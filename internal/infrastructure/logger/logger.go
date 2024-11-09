package logger

import (
	"context"
)

func Info(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Info(args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	fromContext(ctx).Infof(template, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Warn(args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	fromContext(ctx).Warnf(template, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	fromContext(ctx).Error(args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	fromContext(ctx).Errorf(template, args...)
}
