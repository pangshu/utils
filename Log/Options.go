package Log

import (
    "os"
    "strings"
    "sync"
    "time"
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

func WithLevel(typeValue string) Option {
    return func(r *RotateLog) {
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

//func getLevel(level string) zapcore.Level {
//	switch level {
//	case "debug":
//		return zapcore.DebugLevel
//	case "info":
//		return zapcore.InfoLevel
//	case "warn":
//		return zapcore.WarnLevel
//	case "error":
//		return zapcore.ErrorLevel
//	case "dpanic":
//		return zapcore.DPanicLevel
//	case "panic":
//		return zapcore.PanicLevel
//	case "fatal":
//		return zapcore.FatalLevel
//	default:
//		return zapcore.ErrorLevel
//	}
//}
