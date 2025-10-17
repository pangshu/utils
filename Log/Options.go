package Log

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
)

type RotateLog struct {
	FilePath         string `json:"filePath" yaml:"filePath" default:"./log/"`                                    // 日志文件路径
	AppName          string `json:"appName" yaml:"appName" default:"app"`                                         // 日志文件名
	Level            string `json:"level" yaml:"level" default:"all"`                                             // 日志级别，默认为all
	Type             string `json:"type" yaml:"type" default:"all"`                                               // 日志类型，默认为all, all 表示所有 daily 表示按天 size 表示按大小
	RotateTime       int    `json:"rotateTime" yaml:"rotateTime" default:"86400"`                                 // 日志文件切割时间, 单位:秒
	RotateSize       int    `json:"rotateSize" yaml:"rotateSize" default:"100"`                                   // 日志文件切割大小, 单位:MB
	MaxBackups       int    `json:"maxBackups" yaml:"maxBackups" default:"100"`                                   // 日志文件保存数量
	MaxSize          int    `json:"maxSize" yaml:"maxSize" default:"1024"`                                        // 日志文件保存最大值, 单位:MB
	MaxAge           int    `json:"maxAge" yaml:"maxAge" default:"31536000"`                                      // 日志文件保存时间
	LocalTime        bool   `json:"localTime" yaml:"localTime" default:"true"`                                    // 是否使用本地时间
	Compress         bool   `json:"compress" yaml:"compress" default:"false"`                                     // 日志文件是否压缩
	Stdout           bool   `json:"stdout" yaml:"stdout" default:"true"`                                          // 是否输出到控制台
	BackupTimeFormat string `json:"backupTimeFormat" yaml:"backupTimeFormat" default:"2006-01-02T15:04:05Z07:00"` // 日志文件保存时间格式
	// zapcore.EncoderConfig配置项，默认不需要配置
	MessageKey       string `json:"messageKey" yaml:"messageKey"` // 输入信息的key名
	LevelKey         string `json:"levelKey" yaml:"levelKey"`     // 输出日志级别的key名
	TimeKey          string `json:"timeKey" yaml:"timeKey"`       // 输出时间的key名
	NameKey          string `json:"nameKey" yaml:"nameKey"`
	CallerKey        string `json:"callerKey" yaml:"callerKey"`         //输出调用者的key名
	FunctionKey      string `json:"functionKey" yaml:"functionKey"`     //输出函数的key名
	StacktraceKey    string `json:"stacktraceKey" yaml:"stacktraceKey"` // 输出栈信息的key名
	SkipLineEnding   string `json:"skipLineEnding" yaml:"skipLineEnding"`
	LineEnding       string `json:"lineEnding" yaml:"lineEnding"`           // 每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
	EncodeLevel      string `json:"levelEncoder" yaml:"levelEncoder"`       // 基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
	EncodeTime       string `json:"timeEncoder" yaml:"timeEncoder"`         // 输出的时间格式
	EncodeDuration   string `json:"durationEncoder" yaml:"durationEncoder"` //一般zapcore.SecondsDurationEncoder,执行消耗时间转化成浮点型的秒
	EncodeCaller     string `json:"callerEncoder" yaml:"callerEncoder"`     //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
	EncodeName       string `json:"nameEncoder" yaml:"nameEncoder"`
	ConsoleSeparator string `json:"consoleSeparator" yaml:"consoleSeparator"`

	EncoderConfig zapcore.EncoderConfig `json:"encoderConfig" yaml:"encoderConfig"`

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

type EncoderConfig struct {
}

var (
	currentTime = time.Now
	osStat      = os.Stat
	megabyte    = 1024 * 1024
)

const (
	backupTimeFormat = "2006-01-02T15:04:05Z07:00"
	compressSuffix   = ".gz"
	defaultSuffix    = ".log"
	defaultMaxSize   = 100
	defaultRole      = all
)

const (
	daily = 1
	size  = 1 << 1
	all   = 1 << 2
)

type Option func(*RotateLog)

// 设置默认值
//func (l *Log) DefaultOptions() func(*RotateLog) {
//	// TODO 后续完善
//	return func(r *RotateLog) {
//		l.SetDefaults(r)
//		r.EncoderConfig = zap.NewProductionEncoderConfig()
//	}
//}

func (l *Log) WithLevel(level string) Option {

	fmt.Println("++++++>>>>>>++++++")
	return func(r *RotateLog) {
		fmt.Println("++++++++++++")
		fmt.Println("level:", level)
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

func (l *Log) WithType(typeValue string) Option {
	return func(r *RotateLog) {
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

func (l *Log) WithTypeByInt(typeValue int) Option {
	return func(r *RotateLog) {
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

func (l *Log) WithRotateTime(duration int) Option {
	return func(r *RotateLog) {
		r.RotateTime = duration
	}
}

func (l *Log) WithRotateSize(size int) Option {
	return func(r *RotateLog) {
		r.RotateSize = size
	}
}

func (l *Log) WithMaxBackups(max int) Option {
	return func(r *RotateLog) {
		r.MaxBackups = max
	}
}

func (l *Log) WithMaxSize(size int) Option {
	return func(r *RotateLog) {
		r.MaxSize = size
	}
}

func (l *Log) WithMaxAge(duration int) Option {
	return func(r *RotateLog) {
		r.MaxAge = duration
	}
}

func (l *Log) WithLocalTime(localTime bool) Option {
	return func(r *RotateLog) {
		r.LocalTime = localTime
	}
}

func (l *Log) WithCompress(compress bool) Option {
	return func(r *RotateLog) {
		r.Compress = compress
	}
}

func (l *Log) WithBackupTimeFormat(tpl string) Option {
	return func(r *RotateLog) {
		r.BackupTimeFormat = tpl
	}
}

// 配置EncoderConfig各项值

func (l *Log) WithMessageKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.MessageKey = v
	}
}

func (l *Log) WithLevelKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.LevelKey = v
	}
}
func (l *Log) WithTimeKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.TimeKey = v
	}
}
func (l *Log) WithNameKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.NameKey = v
	}
}
func (l *Log) WithCallerKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.CallerKey = v
	}
}
func (l *Log) WithFunctionKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.FunctionKey = v
	}
}
func (l *Log) WithStacktraceKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.StacktraceKey = v
	}
}
func (l *Log) WithSkipLineEnding(v bool) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.SkipLineEnding = v
	}
}
func (l *Log) WithLineEnding(v string) Option {
	return func(r *RotateLog) {
		if len(v) > 0 {
			r.EncoderConfig.LineEnding = v
		} else {
			r.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
		}
	}
}

func (l *Log) WithConsoleSeparator(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.ConsoleSeparator = v
	}
}

func (l *Log) WithEncodeLevel(level string) Option {
	return func(r *RotateLog) {
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

func (l *Log) WithEncodeTime(level string, layout ...string) Option {
	return func(r *RotateLog) {
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

func (l *Log) WithEncodeDuration(level string) Option {
	return func(r *RotateLog) {
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

func (l *Log) WithEncodeCaller(level string) Option {
	return func(r *RotateLog) {
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

func (l *Log) WithEncodeName(level string) Option {
	return func(r *RotateLog) {
		switch strings.ToLower(level) {
		case strings.ToLower("Full"):
			r.EncoderConfig.EncodeName = zapcore.FullNameEncoder
		default:
			r.EncoderConfig.EncodeName = zapcore.FullNameEncoder
		}
	}
}

func (l *Log) WithStdout(typeValue bool) Option {
	return func(r *RotateLog) {
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

//
//func setDefaults(p *RotateLog) {
//	// Iterate over the fields of the Person struct using reflection
//	// and set the default value for each field if the field is not provided
//	// by the caller of the constructor function.
//	for i := 0; i < reflect.TypeOf(*p).NumField(); i++ {
//		field := reflect.TypeOf(*p).Field(i)
//		if value, ok := field.Tag.Lookup("default"); ok {
//			switch field.Type.Kind() {
//			case reflect.String:
//				if p.Name == "" {
//					p.Name = value
//				}
//			case reflect.Int:
//				if p.Age == 0 {
//					if intValue, err := strconv.Atoi(value); err == nil {
//						p.Age = intValue
//					}
//				}
//			}
//		}
//	}
//}

func (l *Log) SetDefaults(cfg *RotateLog) {
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
