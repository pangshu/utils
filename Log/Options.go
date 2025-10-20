package Log

import (
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level
type Field = zap.Field
type Option = zap.Option
type EncoderConfig = zapcore.EncoderConfig

var (
	WrapCore      = zap.WrapCore
	Hooks         = zap.Hooks
	Fields        = zap.Fields
	ErrorOutput   = zap.ErrorOutput
	Development   = zap.Development
	AddCaller     = zap.AddCaller
	WithCaller    = zap.WithCaller
	AddCallerSkip = zap.AddCallerSkip
	AddStacktrace = zap.AddStacktrace
	IncreaseLevel = zap.IncreaseLevel
	WithPanicHook = zap.WithPanicHook
	WithFatalHook = zap.WithFatalHook
	WithClock     = zap.WithClock
)

type RotateConfig struct {
	FilePath         string        `json:"filePath" yaml:"filePath" default:"./log/"`                                    // 日志文件路径
	AppName          string        `json:"appName" yaml:"appName" default:"app"`                                         // 日志文件名
	Level            string        `json:"level" yaml:"level" default:"all"`                                             // 日志级别，默认为all
	Type             string        `json:"type" yaml:"type" default:"all"`                                               // 日志类型，默认为all, all 表示所有 daily 表示按天 size 表示按大小
	RotateTime       int           `json:"rotateTime" yaml:"rotateTime" default:"86400"`                                 // 日志文件切割时间, 单位:秒
	RotateSize       int           `json:"rotateSize" yaml:"rotateSize" default:"100"`                                   // 日志文件切割大小, 单位:MB
	MaxBackups       int           `json:"maxBackups" yaml:"maxBackups" default:"100"`                                   // 日志文件保存数量
	MaxSize          int           `json:"maxSize" yaml:"maxSize" default:"1024"`                                        // 日志文件保存最大值, 单位:MB
	MaxAge           int           `json:"maxAge" yaml:"maxAge" default:"31536000"`                                      // 日志文件保存时间
	LocalTime        bool          `json:"localTime" yaml:"localTime" default:"true"`                                    // 是否使用本地时间
	Compress         bool          `json:"compress" yaml:"compress" default:"false"`                                     // 日志文件是否压缩
	Stdout           bool          `json:"stdout" yaml:"stdout" default:"true"`                                          // 是否输出到控制台
	BackupTimeFormat string        `json:"backupTimeFormat" yaml:"backupTimeFormat" default:"2006-01-02T15:04:05Z07:00"` // 日志文件保存时间格式
	EncoderConfig    EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`

	stdout         io.Writer
	level          zapcore.Level
	file           *os.File
	size           int64 // 内容长度
	role           int   // 当前类型权限值
	mutex          *sync.Mutex
	rotateTimeChan <-chan time.Time // notify rotate event
	rotateSizeChan <-chan bool
	close          chan struct{} // close file and write goroutine
}

var (
	currentTime = time.Now
	osStat      = os.Stat
	megabyte    = 1024 * 1024
)

var (
	backupTimeFormat = "2006-01-02T15-04-05.000"
	compressSuffix   = ".gz"
	defaultSuffix    = ".log"
	defaultMaxSize   = 100
	defaultRole      = all
	//defaultEncoderConfig = zap.NewProductionEncoderConfig()
)

const (
	daily = 1
	size  = 1 << 1
	all   = 1 << 2
)

type RotateOption func(*RotateConfig)

// Init 初始化配置
func Init(cfg *RotateConfig, opts ...RotateOption) *RotateConfig {
	if cfg == nil {
		cfg = DefaultRotateOptions()
	} else {
		if cfg.EncoderConfig.MessageKey == "" {
			cfg.EncoderConfig = zap.NewProductionEncoderConfig()
			cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
		}
	}
	//fmt.Println("====", cfg.Stdout)
	//if cfg.Stdout == true {
	//	cfg.stdout = os.Stdout
	//} else {
	//	cfg.stdout = os.Stderr
	//}
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

// DefaultRotateOptions 读取默认配置
func DefaultRotateOptions() *RotateConfig {
	var cfg RotateConfig
	SetDefaults(&cfg)
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	return &cfg
}

func WithLevel(level string) RotateOption {
	return func(r *RotateConfig) {
		switch level {
		case "debug":
			r.level = zapcore.DebugLevel
		case "info":
			r.level = zapcore.InfoLevel
		case "warn":
			r.level = zapcore.WarnLevel
		case "error":
			r.level = zapcore.ErrorLevel
		case "dpanic":
			r.level = zapcore.DPanicLevel
		case "panic":
			r.level = zapcore.PanicLevel
		case "fatal":
			r.level = zapcore.FatalLevel
		default:
			r.level = zapcore.ErrorLevel
		}
	}

}

func WithType(typeValue string) RotateOption {
	return func(r *RotateConfig) {
		switch strings.ToLower(typeValue) {
		case "all":
			r.role = all
		case "daily":
			r.role = daily
		case "size":
			r.role = size
		default:
			r.role = all
		}
	}
}

func WithTypeByInt(typeValue int) RotateOption {
	return func(r *RotateConfig) {
		switch typeValue {
		case 0:
			r.role = all
		case 1:
			r.role = daily
		case 2:
			r.role = size
		default:
			r.role = all
		}
	}
}

func WithRotateTime(duration int) RotateOption {
	return func(r *RotateConfig) {
		r.RotateTime = duration
	}
}

func WithRotateSize(size int) RotateOption {
	return func(r *RotateConfig) {
		r.RotateSize = size
	}
}

func WithMaxBackups(max int) RotateOption {
	return func(r *RotateConfig) {
		r.MaxBackups = max
	}
}

func WithMaxSize(size int) RotateOption {
	return func(r *RotateConfig) {
		r.MaxSize = size
	}
}

func WithMaxAge(duration int) RotateOption {
	return func(r *RotateConfig) {
		r.MaxAge = duration
	}
}

func WithLocalTime(localTime bool) RotateOption {
	return func(r *RotateConfig) {
		r.LocalTime = localTime
	}
}

func WithCompress(compress bool) RotateOption {
	return func(r *RotateConfig) {
		r.Compress = compress
	}
}

func WithBackupTimeFormat(tpl string) RotateOption {
	return func(r *RotateConfig) {
		r.BackupTimeFormat = tpl
	}
}

// 配置EncoderConfig各项值

func WithMessageKey(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.MessageKey = v
	}
}

func WithLevelKey(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.LevelKey = v
	}
}
func WithTimeKey(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.TimeKey = v
	}
}
func WithNameKey(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.NameKey = v
	}
}
func WithCallerKey(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.CallerKey = v
	}
}
func WithFunctionKey(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.FunctionKey = v
	}
}
func WithStacktraceKey(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.StacktraceKey = v
	}
}
func WithSkipLineEnding(v bool) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.SkipLineEnding = v
	}
}
func WithLineEnding(v string) RotateOption {
	return func(r *RotateConfig) {
		if len(v) > 0 {
			r.EncoderConfig.LineEnding = v
		} else {
			r.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
		}
	}
}

func WithConsoleSeparator(v string) RotateOption {
	return func(r *RotateConfig) {
		r.EncoderConfig.ConsoleSeparator = v
	}
}

func WithEncodeLevel(level string) RotateOption {
	return func(r *RotateConfig) {
		switch strings.ToLower(level) {
		case strings.ToLower("LowercaseLevelEncoder"), strings.ToLower("Lowercase"):
			r.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		case strings.ToLower("LowercaseColorLevelEncoder"), strings.ToLower("LowercaseColor"):
			r.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		case strings.ToLower("CapitalLevelEncoder"), strings.ToLower("Capital"):
			r.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		case strings.ToLower("CapitalColorLevelEncoder"), strings.ToLower("CapitalColor"):
			r.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		default:
			r.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		}
	}
}

func WithEncodeTime(level string, layout ...string) RotateOption {
	return func(r *RotateConfig) {
		if strings.ToLower(level) == "time" {
			if len(layout) == 0 {
				r.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
			} else {
				r.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(layout[0])
			}
		} else {
			switch strings.ToLower(level) {
			case strings.ToLower("EpochTimeEncoder"), strings.ToLower("Epoch"):
				r.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder
			case strings.ToLower("EpochMillisTimeEncoder"), strings.ToLower("EpochMillis"):
				r.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
			case strings.ToLower("EpochNanosTimeEncoder"), strings.ToLower("EpochNanos"):
				r.EncoderConfig.EncodeTime = zapcore.EpochNanosTimeEncoder
			case strings.ToLower("ISO8601TimeEncoder"), strings.ToLower("ISO8601"):
				r.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			case strings.ToLower("RFC3339TimeEncoder"), strings.ToLower("RFC3339"):
				r.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
			case strings.ToLower("RFC3339NanoTimeEncoder"), strings.ToLower("RFC3339Nano"):
				r.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
			default:
				r.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
			}
		}
	}
}

func WithEncodeDuration(level string) RotateOption {
	return func(r *RotateConfig) {
		switch strings.ToLower(level) {
		case strings.ToLower("SecondsDurationEncoder"), strings.ToLower("Second"):
			r.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		case strings.ToLower("NanosDurationEncoder"), strings.ToLower("Nano"):
			r.EncoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
		case strings.ToLower("MillisDurationEncoder"), strings.ToLower("Milli"):
			r.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
		case strings.ToLower("StringDurationEncoder"), strings.ToLower("String"):
			r.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
		default:
			r.EncoderConfig.EncodeDuration = zapcore.NanosDurationEncoder
		}
	}
}

func WithEncodeCaller(level string) RotateOption {
	return func(r *RotateConfig) {
		switch strings.ToLower(level) {
		case strings.ToLower("FullCallerEncoder"), strings.ToLower("Full"):
			r.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		case strings.ToLower("ShortCallerEncoder"), strings.ToLower("Short"):
			r.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		default:
			r.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
		}
	}
}

func WithEncodeName(level string) RotateOption {
	return func(r *RotateConfig) {
		switch strings.ToLower(level) {
		case strings.ToLower("Full"):
			r.EncoderConfig.EncodeName = zapcore.FullNameEncoder
		default:
			r.EncoderConfig.EncodeName = zapcore.FullNameEncoder
		}
	}
}

func WithStdout(typeValue bool) RotateOption {
	return func(r *RotateConfig) {
		switch typeValue {
		case true:
			r.stdout = os.Stdout
		case false:
			r.stdout = os.Stderr
		default:
			r.stdout = os.Stdout
		}
	}
}

func SetDefaults(cfg *RotateConfig) {
	v := reflect.ValueOf(cfg)

	// 确保传入的是指针
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return
	}

	// 获取指针指向的结构体
	v = v.Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if !field.CanSet() {
			continue // 字段不可写（非导出字段）
		}

		tag := t.Field(i).Tag.Get("default")
		if tag == "" {
			continue // 没有默认值标签
		}

		// 如果字段是零值，则设置默认值
		if field.Interface() == reflect.Zero(field.Type()).Interface() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(tag)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if val, err := strconv.Atoi(tag); err == nil {
					field.SetInt(int64(val))
				}
			case reflect.Bool:
				boolValue, err := strconv.ParseBool(tag)
				if err == nil {
					field.SetBool(boolValue)
				}
				// 可扩展其他类型：bool、float 等
			default:

			}

		}
	}
}
