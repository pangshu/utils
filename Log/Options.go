package Log

import (
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
)

type RotateLog struct {
	FilePath         string        `json:"file_path" yaml:"file_path"`                   // 日志文件路径
	FileName         string        `json:"file_name" yaml:"file_name"`                   // 日志文件名
	Level            string        `json:"level" yaml:"level"`                           // 日志级别，默认为all
	Type             string        `json:"type" yaml:"type"`                             // 日志类型，默认为all, all 表示所有 daily 表示按天 size 表示按大小
	RotateTime       time.Duration `json:"rotate_time" yaml:"rotate_time"`               // 日志文件切割时间, 单位:秒
	RotateSize       int           `json:"rotate_size" yaml:"rotate_size"`               // 日志文件切割大小, 单位:MB
	MaxBackups       int           `json:"max_backups" yaml:"max_backups"`               // 日志文件保存数量
	MaxSize          int           `json:"max_size" yaml:"max_size"`                     // 日志文件保存最大值, 单位:MB
	MaxAge           time.Duration `json:"max_age" yaml:"max_age"`                       // 日志文件保存时间
	LocalTime        bool          `json:"local_time" yaml:"local_time"`                 // 是否使用本地时间
	Compress         bool          `json:"compress" yaml:"compress"`                     // 日志文件是否压缩
	Stdout           bool          `json:"stdout" yaml:"stdout"`                         // 是否输出到控制台
	BackupTimeFormat string        `json:"backup_time_format" yaml:"backup_time_format"` // 日志文件保存时间格式

	//MessageKey     string                  `json:"message_key" yaml:"message_key"` // 输入信息的key名
	//LevelKey       string                  `json:"level_key" yaml:"level_key"`     // 输出日志级别的key名
	//TimeKey        string                  `json:"time_key" yaml:"time_key"`       // 输出时间的key名
	//NameKey        string                  `json:"name_key" yaml:"name_key"`
	//CallerKey      string                  `json:"caller_key" yaml:"caller_key"`         //输出调用者的key名
	//FunctionKey    string                  `json:"function_key" yaml:"function_key"`     //输出函数的key名
	//StacktraceKey  string                  `json:"stacktrace_key" yaml:"stacktrace_key"` // 输出栈信息的key名
	//SkipLineEnding bool                    `json:"skip_line_ending" yaml:"skip_line_ending"`
	//LineEnding     string                  `json:"line_ending" yaml:"line_ending"`           // 每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
	//EncodeLevel    zapcore.LevelEncoder    `json:"level_encoder" yaml:"level_encoder"`       // 基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
	//EncodeTime     zapcore.TimeEncoder     `json:"time_encoder" yaml:"time_encoder"`         // 输出的时间格式
	//EncodeDuration zapcore.DurationEncoder `json:"duration_encoder" yaml:"duration_encoder"` //一般zapcore.SecondsDurationEncoder,执行消耗时间转化成浮点型的秒
	//EncodeCaller   zapcore.CallerEncoder   `json:"caller_encoder" yaml:"caller_encoder"`     //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
	//EncodeName zapcore.NameEncoder `json:"nameEncoder" yaml:"nameEncoder"`
	EncoderConfig zapcore.EncoderConfig `json:"encoder_config" yaml:"encoder_config"`

	level          zapcore.Level
	file           *os.File
	size           int64 // 内容长度
	role           int   // 当前类型权限值
	mutex          *sync.Mutex
	rotateTimeChan <-chan time.Time // notify rotate event
	rotateSizeChan <-chan bool
	close          chan struct{} // close file and write goroutine
}

//type EncoderConfig struct {
//	MessageKey    string `json:"messageKey" yaml:"messageKey"`
//	LevelKey      string `json:"levelKey" yaml:"levelKey"`
//	TimeKey       string `json:"timeKey" yaml:"timeKey"`
//	NameKey       string `json:"nameKey" yaml:"nameKey"`
//	CallerKey     string `json:"callerKey" yaml:"callerKey"`
//	FunctionKey   string `json:"functionKey" yaml:"functionKey"`
//	StacktraceKey string `json:"stacktraceKey" yaml:"stacktraceKey"`
//	LineEnding    string `json:"lineEnding" yaml:"lineEnding"`
//
//	MessageKey     string          `json:"messageKey" yaml:"messageKey"`       // 输入信息的key名
//	LevelKey       string          `json:"levelKey" yaml:"levelKey"`           // 输出日志级别的key名
//	TimeKey        string          `json:"timeKey" yaml:"timeKey"`             // 输出时间的key名
//	NameKey        string          `json:"nameKey" yaml:"nameKey"`             //
//	CallerKey      string          `json:"callerKey" yaml:"callerKey"`         //输出调用者的key名
//	FunctionKey    string          `json:"functionKey" yaml:"functionKey"`     //输出函数的key名
//	StacktraceKey  string          `json:"stacktraceKey" yaml:"stacktraceKey"` // 输出栈信息的key名
//	SkipLineEnding bool            `json:"skipLineEnding" yaml:"skipLineEnding"`
//	LineEnding     string          `json:"lineEnding" yaml:"lineEnding"`           // 每行的分隔符。基本zapcore.DefaultLineEnding 即"\n"
//	EncodeLevel    LevelEncoder    `json:"levelEncoder" yaml:"levelEncoder"`       // 基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
//	EncodeTime     TimeEncoder     `json:"timeEncoder" yaml:"timeEncoder"`         // 输出的时间格式
//	EncodeDuration DurationEncoder `json:"durationEncoder" yaml:"durationEncoder"` //一般zapcore.SecondsDurationEncoder,执行消耗时间转化成浮点型的秒
//	EncodeCaller   CallerEncoder   `json:"callerEncoder" yaml:"callerEncoder"`     //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
//
//	EncodeName          NameEncoder                      `json:"nameEncoder" yaml:"nameEncoder"`
//	NewReflectedEncoder func(io.Writer) ReflectedEncoder `json:"-" yaml:"-"`
//	ConsoleSeparator    string                           `json:"consoleSeparator" yaml:"consoleSeparator"`
//}

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

func WithEncoderMessageKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.MessageKey = v
	}
}

func WithEncoderLevelKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.LevelKey = v
	}
}
func WithEncoderTimeKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.TimeKey = v
	}
}
func WithEncoderNameKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.NameKey = v
	}
}
func WithEncoderCallerKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.CallerKey = v
	}
}
func WithEncoderFunctionKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.FunctionKey = v
	}
}
func WithEncoderStacktraceKey(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.StacktraceKey = v
	}
}
func WithEncoderSkipLineEnding(v bool) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.SkipLineEnding = v
	}
}
func WithEncoderLineEnding(v string) Option {
	return func(r *RotateLog) {
		r.EncoderConfig.LineEnding = v
	}
}

func WithEncoderEncodeLevel(level string) Option {
	return func(r *RotateLog) {
		switch level {
		case "debug":
			r.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		case "info":
			r.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
		case "warn":
			r.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		case "error":
			r.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		default:
			r.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		}
	}
}

func WithEncoderEncodeLevel(level string) Option {
	return func(r *RotateLog) {
		switch level {
		case "debug":
			r.EncoderConfig.EncodeLevel = zapcore.EpochTimeEncoder
		case "info":
			r.EncoderConfig.EncodeLevel = zapcore.EpochMillisTimeEncoder
		case "warn":
			r.EncoderConfig.EncodeLevel = zapcore.EpochNanosTimeEncoder
		case "error":
			r.EncoderConfig.EncodeLevel = zapcore.ISO8601TimeEncoder
		default:
			r.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		}
	}
}



	EpochTimeEncoder
	EpochMillisTimeEncoder
	EpochNanosTimeEncoder
	ISO8601TimeEncoder
	RFC3339TimeEncoder
	RFC3339NanoTimeEncoder
	TimeEncoderOfLayout

	SecondsDurationEncoder
	NanosDurationEncoder
	MillisDurationEncoder
	StringDurationEncoder

	FullCallerEncoder
	ShortCallerEncoder

	FullNameEncoder

	EncodeLevel    zapcore.LevelEncoder    `json:"level_encoder" yaml:"level_encoder"`       // 基本zapcore.LowercaseLevelEncoder。将日志级别字符串转化为小写
	EncodeTime     zapcore.TimeEncoder     `json:"time_encoder" yaml:"time_encoder"`         // 输出的时间格式
	EncodeDuration zapcore.DurationEncoder `json:"duration_encoder" yaml:"duration_encoder"` //一般zapcore.SecondsDurationEncoder,执行消耗时间转化成浮点型的秒
	EncodeCaller   zapcore.CallerEncoder   `json:"caller_encoder" yaml:"caller_encoder"`     //一般zapcore.ShortCallerEncoder，以包/文件:行号 格式化调用堆栈
}
