package Log

import (
	"os"
	"sync"
	"time"
)

type RotateLog struct {
	FilePath   string        // 日志文件路径
	FileName   string        // 日志文件名
	Level      string        // 日志级别
	LogType    string        // 日志类型，默认为all, all 表示所有 daily 表示按天 size 表示按大小
	RotateTime time.Duration // 日志文件切割时间, 单位:秒
	RotateSize int           // 日志文件切割大小, 单位:MB
	MaxBackups int           // 日志文件保存数量
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

func WithMaxAge(duration time.Duration) Option {
	return func(r *RotateLog) {
		r.MaxAge = duration
	}
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.ErrorLevel
	}
}
