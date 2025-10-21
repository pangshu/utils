package Log

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	l *zap.Logger
	// https://pkg.go.dev/go.uber.org/zap#example-AtomicLevel
	al *zap.AtomicLevel
}

func New(cfg *RotateConfig, opts ...Option) *Logger {
	if cfg == nil {
		cfg = Init(cfg)
	}

	var syncer = []zapcore.WriteSyncer{
		zapcore.AddSync(rotateWriter(cfg)),
		zapcore.AddSync(cfg.stdout),
	}
	//syncer = append(syncer, zapcore.AddSync(rotateWriter(cfg)))
	//syncer = append(syncer, zapcore.AddSync(cfg.stdout))

	al := zap.NewAtomicLevelAt(cfg.level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(syncer...),
		al,
	)

	//zapcore.NewCore(encoder(encoderConfig), zapcore.NewMultiWriteSyncer(wsInfo...), infoLevel()),
	return &Logger{l: zap.New(core, opts...), al: &al}
}

func rotateWriter(cfg *RotateConfig) io.Writer {
	r, err := NewRotate(cfg)
	if err != nil {
		panic(err)
	}
	return r
}

func (l *Logger) SetLevel(level Level) {
	// 使用Tee模式本身已经根据不同日志级别
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
