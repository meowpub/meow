package clog

import (
	"context"

	"github.com/meowpub/meow/lib"
	"go.uber.org/zap"
)

func DPanic(ctx context.Context, msg string, fields ...zap.Field) {
	lib.GetLogger(ctx).DPanic(msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	lib.GetLogger(ctx).Debug(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	lib.GetLogger(ctx).Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	lib.GetLogger(ctx).Fatal(msg, fields...)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	lib.GetLogger(ctx).Info(msg, fields...)
}

func Panic(ctx context.Context, msg string, fields ...zap.Field) {
	lib.GetLogger(ctx).Panic(msg, fields...)
}

func Warn(ctx context.Context, msg string, fields ...zap.Field) {
	lib.GetLogger(ctx).Warn(msg, fields...)
}
