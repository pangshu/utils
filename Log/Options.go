package Log

import (
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
)

type RotateLog struct {
	FilePath         string        `json:"filePath" yaml:"filePath"`                 // 日志文件路径
	AppName          string        `json:"appName" yaml:"appName"`                   // 日志文件名
	Level            string        `json:"level" yaml:"level"`                       // 日志级别，默认为all
	Type             string        `json:"type" yaml:"type"`                         // 日志类型，默认为all, all 表示所有 daily 表示按天 size 表示按大小
	RotateTime       time.Duration `json:"rotateTime" yaml:"rotateTime"`             // 日志文件切割时间, 单位:秒
	RotateSize       int           `json:"rotateSize" yaml:"rotateSize"`             // 日志文件切割大小, 单位:MB
	MaxBackups       int           `json:"maxBackups" yaml:"maxBackups"`             // 日志文件保存数量
	MaxSize          int           `json:"maxSize" yaml:"maxSize"`                   // 日志文件保存最大值, 单位:MB
	MaxAge           time.Duration `json:"maxAge" yaml:"maxAge"`                     // 日志文件保存时间
	LocalTime        bool          `json:"localTime" yaml:"localTime"`               // 是否使用本地时间
	Compress         bool          `json:"compress" yaml:"compress"`                 // 日志文件是否压缩
	Stdout           bool          `json:"stdout" yaml:"stdout"`                     // 是否输出到控制台
	BackupTimeFormat string        `json:"backupTimeFormat" yaml:"backupTimeFormat"` // 日志文件保存时间格式
	// zapcore.EncoderConfig配置项，默认不需要配置
	MessageKey       string                  `json:"messageKey" yaml:"messageKey"` // 输入信息的key名
	LevelKey         string                  `json:"levelKey" yaml:"levelKey"`     // 输出日志级别的key名
	TimeKey          string                  `json:"timeKey" yaml:"timeKey"`       // 输出时间的key名
	NameKey          string                  `json:"nameKey" yaml:"nameKey"`
	CallerKey        string                  `json:"callerKey" yaml:"callerKey"`         //输出调用者的key名
	FunctionKey      string                  `json:"functionKey" yaml:"functionKey"`     //输出函数的key名
	StacktraceKey    string                  `json:"stacktraceKey" yaml:"stacktraceKey"` // 输出栈信息的key名
	SkipLineEnding   bool                    `json:"skipLineEnding" yaml:"skipLineEnding"`
	LineEnding       string                  `json:"lineEnding" yaml:"lineEnding"`           // 每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
	EncodeLevel      zapcore.LevelEncoder    `json:"levelEncoder" yaml:"levelEncoder"`       // 基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
	EncodeTime       zapcore.TimeEncoder     `json:"timeEncoder" yaml:"timeEncoder"`         // 输出的时间格式
	EncodeDuration   zapcore.DurationEncoder `json:"durationEncoder" yaml:"durationEncoder"` //一般zapcore.SecondsDurationEncoder,执行消耗时间转化成浮点型的秒
	EncodeCaller     zapcore.CallerEncoder   `json:"callerEncoder" yaml:"callerEncoder"`     //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
	EncodeName       zapcore.NameEncoder     `json:"nameEncoder" yaml:"nameEncoder"`
	ConsoleSeparator string                  `json:"consoleSeparator" yaml:"consoleSeparator"`

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

const (
	backupTimeFormat = "2006-01-02T15-04-05.000"
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
func defaultOptions() Option {
	// TODO 后续完善
	return func(r *RotateLog) {
	}
}

func WithLevel(level string) Option {
	return func(r *RotateLog) {
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

func WithType(typeValue string) Option {
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

func WithTypeByInt(typeValue int) Option {
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

func WithRotateTime(duration time.Duration) Option {
	return func(r *RotateLog) {
		r.RotateTime = duration
	}
}

func WithRotateSize(size int) Option {
	return func(r *RotateLog) {
		r.RotateSize = size
	}
}

func WithMaxBackups(max int) Option {
	return func(r *RotateLog) {
		r.MaxBackups = max
	}
}

func WithMaxSize(size int) Option {
	return func(r *RotateLog) {
		r.MaxSize = size
	}
}

func WithMaxAge(duration time.Duration) Option {
	return func(r *RotateLog) {
		r.MaxAge = duration
	}
}

func WithLocalTime(localTime bool) Option {
	return func(r *RotateLog) {
		r.LocalTime = localTime
	}
}

func WithCompress(compress bool) Option {
	return func(r *RotateLog) {
		r.Compress = compress
	}
}

func WithBackupTimeFormat(tpl string) Option {
	return func(r *RotateLog) {
		r.BackupTimeFormat = tpl
	}
}

// 配置EncoderConfig各项值

func WithMessageKey(v string) Option {
	return func(r *RotateLog) {
		r.MessageKey = v
	}
}

func WithLevelKey(v string) Option {
	return func(r *RotateLog) {
		r.LevelKey = v
	}
}
func WithTimeKey(v string) Option {
	return func(r *RotateLog) {
		r.TimeKey = v
	}
}
func WithNameKey(v string) Option {
	return func(r *RotateLog) {
		r.NameKey = v
	}
}
func WithCallerKey(v string) Option {
	return func(r *RotateLog) {
		r.CallerKey = v
	}
}
func WithFunctionKey(v string) Option {
	return func(r *RotateLog) {
		r.FunctionKey = v
	}
}
func WithStacktraceKey(v string) Option {
	return func(r *RotateLog) {
		r.StacktraceKey = v
	}
}
func WithSkipLineEnding(v bool) Option {
	return func(r *RotateLog) {
		r.SkipLineEnding = v
	}
}
func WithLineEnding(v string) Option {
	return func(r *RotateLog) {
		if len(v) > 0 {
			r.LineEnding = v
		} else {
			r.LineEnding = zapcore.DefaultLineEnding
		}
	}
}

func WithConsoleSeparator(v string) Option {
	return func(r *RotateLog) {
		r.ConsoleSeparator = v
	}
}

func WithEncodeLevel(level string) Option {
	return func(r *RotateLog) {
		switch strings.ToLower(level) {
		case strings.ToLower("LowercaseLevelEncoder"), strings.ToLower("Lowercase"):
			r.EncodeLevel = zapcore.LowercaseLevelEncoder
		case strings.ToLower("LowercaseColorLevelEncoder"), strings.ToLower("LowercaseColor"):
			r.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		case strings.ToLower("CapitalLevelEncoder"), strings.ToLower("Capital"):
			r.EncodeLevel = zapcore.CapitalLevelEncoder
		case strings.ToLower("CapitalColorLevelEncoder"), strings.ToLower("CapitalColor"):
			r.EncodeLevel = zapcore.CapitalColorLevelEncoder
		default:
			r.EncodeLevel = zapcore.LowercaseLevelEncoder
		}
	}
}

func WithEncodeTime(level string, layout ...string) Option {
	return func(r *RotateLog) {
		if strings.ToLower(level) == "time" {
			if len(layout) == 0 {
				r.EncodeTime = zapcore.RFC3339TimeEncoder
			} else {
				r.EncodeTime = zapcore.TimeEncoderOfLayout(layout[0])
			}
		} else {
			switch strings.ToLower(level) {
			case strings.ToLower("EpochTimeEncoder"), strings.ToLower("Epoch"):
				r.EncodeTime = zapcore.EpochTimeEncoder
			case strings.ToLower("EpochMillisTimeEncoder"), strings.ToLower("EpochMillis"):
				r.EncodeTime = zapcore.EpochMillisTimeEncoder
			case strings.ToLower("EpochNanosTimeEncoder"), strings.ToLower("EpochNanos"):
				r.EncodeTime = zapcore.EpochNanosTimeEncoder
			case strings.ToLower("ISO8601TimeEncoder"), strings.ToLower("ISO8601"):
				r.EncodeTime = zapcore.ISO8601TimeEncoder
			case strings.ToLower("RFC3339TimeEncoder"), strings.ToLower("RFC3339"):
				r.EncodeTime = zapcore.RFC3339TimeEncoder
			case strings.ToLower("RFC3339NanoTimeEncoder"), strings.ToLower("RFC3339Nano"):
				r.EncodeTime = zapcore.RFC3339NanoTimeEncoder
			default:
				r.EncodeTime = zapcore.RFC3339TimeEncoder
			}
		}
	}
}

func WithEncodeDuration(level string) Option {
	return func(r *RotateLog) {
		switch strings.ToLower(level) {
		case strings.ToLower("SecondsDurationEncoder"), strings.ToLower("Second"):
			r.EncodeDuration = zapcore.SecondsDurationEncoder
		case strings.ToLower("NanosDurationEncoder"), strings.ToLower("Nano"):
			r.EncodeDuration = zapcore.NanosDurationEncoder
		case strings.ToLower("MillisDurationEncoder"), strings.ToLower("Milli"):
			r.EncodeDuration = zapcore.MillisDurationEncoder
		case strings.ToLower("StringDurationEncoder"), strings.ToLower("String"):
			r.EncodeDuration = zapcore.StringDurationEncoder
		default:
			r.EncodeDuration = zapcore.NanosDurationEncoder
		}
	}
}

func WithEncodeCaller(level string) Option {
	return func(r *RotateLog) {
		switch strings.ToLower(level) {
		case strings.ToLower("FullCallerEncoder"), strings.ToLower("Full"):
			r.EncodeCaller = zapcore.FullCallerEncoder
		case strings.ToLower("ShortCallerEncoder"), strings.ToLower("Short"):
			r.EncodeCaller = zapcore.ShortCallerEncoder
		default:
			r.EncodeCaller = zapcore.FullCallerEncoder
		}
	}
}

func WithEncodeName(level string) Option {
	return func(r *RotateLog) {
		switch strings.ToLower(level) {
		case strings.ToLower("Full"):
			r.EncodeName = zapcore.FullNameEncoder
		default:
			r.EncodeName = zapcore.FullNameEncoder
		}
	}
}
