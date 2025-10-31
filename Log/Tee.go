package Log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LevelEnablerFunc func(Level) bool
type TeeOption struct {
	Level string `json:"level"`
}

// NewTee 根据日志级别写入多个输出
// https://pkg.go.dev/go.uber.org/zap#example-package-AdvancedConfiguration
func NewTee(cfg *RotateConfig, tees []TeeOption, opts ...Option) *Logger {
	if cfg == nil {
		cfg = Init(cfg)
	}

	var cores []zapcore.Core
	for _, tee := range tees {
		cfg.AppName = fmt.Sprintf("%s-%s", cfg.AppName, tee)
		var syncer = []zapcore.WriteSyncer{
			zapcore.AddSync(rotateWriter(cfg)),
			zapcore.AddSync(cfg.stdout),
		}

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.NewMultiWriteSyncer(syncer...),
			//zap.LevelEnablerFunc(func(level zapcore.Level) bool {
			//	return tee.LevelEnablerFunc(level)
			//}),
			getLevel(tee.Level),
		)
		cores = append(cores, core)
	}
	return &Logger{l: zap.New(zapcore.NewTee(cores...), opts...)}
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
