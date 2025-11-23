package logger

import (
	"context"

	"go.uber.org/zap"
)

type Key struct{}

func NewProductionLogger() (*zap.SugaredLogger, error) {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	logger, err := config.Build()
	if err != nil {
		return nil, err
	}
	return logger.Sugar(), nil
}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, Key{}, logger)
}

func InfoKV(ctx context.Context, msg string, args ...interface{}) {
	logger := getLoggerFromContext(ctx)
	logger.Infow(msg, args...)
}

func ErrorKV(ctx context.Context, format string, args ...interface{}) {
	logger := getLoggerFromContext(ctx)
	logger.Errorw(format, args...)
}

func Fatalf(ctx context.Context, msg string, args ...interface{}) {
	logger := getLoggerFromContext(ctx)
	logger.Fatalf(msg, args...)
}

func getLoggerFromContext(ctx context.Context) *zap.SugaredLogger {
	logger, ok := ctx.Value(Key{}).(*zap.SugaredLogger)
	if !ok {
		return zap.NewNop().Sugar()
	}
	return logger
}
