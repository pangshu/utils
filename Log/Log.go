package Log

import (
	"fmt"

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

type Log struct {
	l *zap.Logger
	// https://pkg.go.dev/go.uber.org/zap#example-AtomicLevel
	al *zap.AtomicLevel
}

func (l *Log) New(conf *RotateLog, opts ...func(*RotateLog)) *Log {
	fmt.Println("=== 开始初始化测试 === ")
	cfg := l.initConfig(conf, opts...)
	fmt.Println(string(cfg.AppName))
	//o := &LogConfig{
	//	Path:       "./log/",
	//	Name:       "output",
	//	Level:      "debug",
	//	MaxSize:    100,
	//	MaxAge:     7,
	//	Stacktrace: "error",
	//	Stdout:     true,
	//}
	//for _, opt := range opts {
	//	opt(conf)
	//}
	//writers := []zapcore.WriteSyncer{newRollingFile(o.Path, o.Name, o.MaxSize, o.MaxAge)}
	//if o.Stdout == true {
	//	writers = append(writers, os.Stdout)
	//}
	//logger := newZapLogger(getLevel(o.Level), getLevel(o.Stacktrace), zapcore.NewMultiWriteSyncer(writers...))
	//zap.RedirectStdLog(logger)

	fmt.Println(">>>>>>>>>>")
	fmt.Println(cfg.Level)
	fmt.Println(">>>>>>>>>>")
	fmt.Println(cfg.level)

	//al := zap.NewAtomicLevelAt(level)
	//core := zapcore.NewCore(
	//	zapcore.NewJSONEncoder(cfg),
	//	zapcore.AddSync(out),
	//	al,
	//)
	//
	//var encoder zapcore.Encoder
	//dyn := zap.NewAtomicLevel()
	////encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	////encoder = zapcore.NewJSONEncoder(encCfg) // zapcore.NewConsoleEncoder(encCfg)
	//dyn.SetLevel(level)
	//encCfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	//encoder = zapcore.NewJSONEncoder(encCfg)
	//
	//return zap.New(zapcore.NewCore(encoder, output, dyn), zap.AddCaller(), zap.AddStacktrace(stacktrace), zap.AddCallerSkip(2))
	//return &Log{l: zap.New()}
	return &Log{}
}

// 初始化配置文件
func (l *Log) initConfig(cfg *RotateLog, opts ...func(*RotateLog)) *RotateLog {
	if cfg == nil {
		cfg = l.getDefaultOptions()
		l.WithLevel(cfg.Level)
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(cfg)
		}
	}
	return cfg
}

// 读取默认配置
func (l *Log) getDefaultOptions() *RotateLog {
	var cfg RotateLog
	l.SetDefaults(&cfg)
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	return &cfg
}

//
//func NewAAA(out io.Writer, level Level, opts ...ZapOption) *Logger {
//	if out == nil {
//		out = os.Stderr
//	}
//
//	al := zap.NewAtomicLevelAt(level)
//	cfg := zap.NewProductionEncoderConfig()
//	cfg.EncodeTime = zapcore.RFC3339TimeEncoder
//
//	core := zapcore.NewCore(
//		zapcore.NewJSONEncoder(cfg),
//		zapcore.AddSync(out),
//		al,
//	)
//	return &Logger{l: zap.New(core, opts...), al: &al}
//}
//
//func (l *Logger) getEncoderConfig(conf RotateLog) zapcore.EncoderConfig {
//	var config zapcore.EncoderConfig
//	if len(conf.MessageKey) > 0 {
//		config.MessageKey = conf.MessageKey
//	}
//
//	return zapcore.EncoderConfig{
//		TimeKey:        "ts",
//		LevelKey:       "level",
//		NameKey:        "logger",
//		CallerKey:      "caller",
//		FunctionKey:    zapcore.OmitKey,
//		MessageKey:     "msg",
//		StacktraceKey:  "stacktrace",
//		LineEnding:     zapcore.DefaultLineEnding,
//		EncodeLevel:    zapcore.LowercaseLevelEncoder,
//		EncodeTime:     zapcore.EpochTimeEncoder,
//		EncodeDuration: zapcore.SecondsDurationEncoder,
//		EncodeCaller:   zapcore.ShortCallerEncoder,
//	}
//}
//
//// SetLevel 动态更改日志级别
//// 对于使用 NewTee 创建的 Logger 无效，因为 NewTee 本意是根据不同日志级别
//// 创建的多个 zap.Core，不应该通过 SetLevel 将多个 zap.Core 日志级别统一
//func (l *Logger) SetLevel(level Level) {
//	if l.al != nil {
//		l.al.SetLevel(level)
//	}
//}
//
//func (l *Logger) Debug(msg string, fields ...Field) {
//	l.l.Debug(msg, fields...)
//}
//
//func (l *Logger) Info(msg string, fields ...Field) {
//	l.l.Info(msg, fields...)
//}
//
//func (l *Logger) Warn(msg string, fields ...Field) {
//	l.l.Warn(msg, fields...)
//}
//
//func (l *Logger) Error(msg string, fields ...Field) {
//	l.l.Error(msg, fields...)
//}
//
//func (l *Logger) Panic(msg string, fields ...Field) {
//	l.l.Panic(msg, fields...)
//}
//
//func (l *Logger) Fatal(msg string, fields ...Field) {
//	l.l.Fatal(msg, fields...)
//}
//
//func (l *Logger) Sync() error {
//	return l.l.Sync()
//}

//var std = New(os.Stderr, InfoLevel)
//
//func Default() *Logger         { return std }
//func ReplaceDefault(l *Logger) { std = l }
//
//func SetLevel(level Level) { std.SetLevel(level) }
//
//func Debug(msg string, fields ...Field) { std.Debug(msg, fields...) }
//func Info(msg string, fields ...Field)  { std.Info(msg, fields...) }
//func Warn(msg string, fields ...Field)  { std.Warn(msg, fields...) }
//func Error(msg string, fields ...Field) { std.Error(msg, fields...) }
//func Panic(msg string, fields ...Field) { std.Panic(msg, fields...) }
//func Fatal(msg string, fields ...Field) { std.Fatal(msg, fields...) }
//
//func Sync() error { return std.Sync() }
