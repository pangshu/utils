package Log

import (
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level = zapcore.Level
type Field = zap.Field
type ZapOption = zap.Option

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

type Logger struct {
	l *zap.Logger
	// https://pkg.go.dev/go.uber.org/zap#example-AtomicLevel
	al *zap.AtomicLevel
}

func (*Log) NewLogger(conf RotateLog, opts ...Option) *zap.Logger {

	o := &LogConfig{
		Path:       "./log/",
		Name:       "output",
		Level:      "debug",
		MaxSize:    100,
		MaxAge:     7,
		Stacktrace: "error",
		Stdout:     true,
	}
	for _, opt := range opts {
		opt(o)
	}
	writers := []zapcore.WriteSyncer{newRollingFile(o.Path, o.Name, o.MaxSize, o.MaxAge)}
	if o.Stdout == true {
		writers = append(writers, os.Stdout)
	}
	logger := newZapLogger(getLevel(o.Level), getLevel(o.Stacktrace), zapcore.NewMultiWriteSyncer(writers...))
	zap.RedirectStdLog(logger)

	return logger
}

func NewAAA(out io.Writer, level Level, opts ...ZapOption) *Logger {
	if out == nil {
		out = os.Stderr
	}

	al := zap.NewAtomicLevelAt(level)
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg),
		zapcore.AddSync(out),
		al,
	)
	return &Logger{l: zap.New(core, opts...), al: &al}
}

func (l *Logger) getEncoderConfig(conf RotateLog) zapcore.EncoderConfig {
	var config zapcore.EncoderConfig
	if len(conf.MessageKey) > 0 {
		config.MessageKey = conf.MessageKey
	}

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

	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}


// SetLevel 动态更改日志级别
// 对于使用 NewTee 创建的 Logger 无效，因为 NewTee 本意是根据不同日志级别
// 创建的多个 zap.Core，不应该通过 SetLevel 将多个 zap.Core 日志级别统一
func (l *Logger) SetLevel(level Level) {
	if l.al != nil {
		l.al.SetLevel(level)
	}
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.l.Sync()
}

var std = New(os.Stderr, InfoLevel)

func Default() *Logger         { return std }
func ReplaceDefault(l *Logger) { std = l }

func SetLevel(level Level) { std.SetLevel(level) }

func Debug(msg string, fields ...Field) { std.Debug(msg, fields...) }
func Info(msg string, fields ...Field)  { std.Info(msg, fields...) }
func Warn(msg string, fields ...Field)  { std.Warn(msg, fields...) }
func Error(msg string, fields ...Field) { std.Error(msg, fields...) }
func Panic(msg string, fields ...Field) { std.Panic(msg, fields...) }
func Fatal(msg string, fields ...Field) { std.Fatal(msg, fields...) }

func Sync() error { return std.Sync() }

func getEncoder() zapcore.EncoderConfig {}
