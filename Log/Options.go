package Log

import (
    "os"
    "strings"
    "sync"
    "time"
)

type RotateLog struct {
    FilePath   string        // 日志文件路径
    FileName   string        // 日志文件名
    Level      string        // 日志级别
    Type       string        // 日志类型，默认为all, all 表示所有 daily 表示按天 size 表示按大小
    RotateTime time.Duration // 日志文件切割时间, 单位:秒
    RotateSize int           // 日志文件切割大小, 单位:MB
    MaxBackups int           // 日志文件保存数量
    MaxSize    int           // 日志文件保存最大值, 单位:MB
    MaxAge     time.Duration // 日志文件保存时间
    LocalTime  bool          // 是否使用本地时间
    Compress   bool          // 日志文件是否压缩
    Stdout     bool

    file   *os.File
    size   int64 // 内容长度
    mutex  *sync.Mutex
    rotate <-chan time.Time // notify rotate event
    close  chan struct{}    // close file and write goroutine
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
            r.Type = "all"
        case "daily":
            r.Type = "daily"
        case "size":
            r.Type = "size"
        default:
            r.Type = "all"
        }
    }
}

func WithTypeByInt(typeValue int) Option {
    return func(r *RotateLog) {
        switch typeValue {
        case 0:
            r.Type = "all"
        case 1:
            r.Type = "daily"
        case 2:
            r.Type = "size"
        default:
            r.Type = "all"
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
