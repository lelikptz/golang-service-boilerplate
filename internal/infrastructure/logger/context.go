package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey struct{}

var loggerContextKey = contextKey{}

func fromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(loggerContextKey).(*zap.SugaredLogger); ok {
		return logger
	}

	return zap.NewNop().Sugar()
}

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}
