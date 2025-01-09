package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	logger      *zap.SugaredLogger
	defFn       func() error
	atomicLevel *zap.AtomicLevel
)

func Default() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}

	env := os.Getenv("APP_ENV")
	al := zap.NewAtomicLevel()
	atomicLevel = &al

	opts := []zap.Option{zap.IncreaseLevel(atomicLevel)}

	var l *zap.Logger
	if env == "prod" {
		l, _ = zap.NewProduction(opts...)
	} else {
		l, _ = zap.NewDevelopment(opts...)
	}
	logger = l.Sugar()
	defFn = logger.Sync

	return logger
}

func CancelLogger() {
	if defFn != nil {
		_ = defFn()
	}
}

func SetLogLevel(level string) {
	if logger == nil {
		Default()
	}

	lvl, err := zapcore.ParseLevel(level)
	if err != nil {
		lvl = zapcore.InfoLevel
	}

	atomicLevel.SetLevel(lvl)
}
