package lib

import (
	"context"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"
	"go.uber.org/zap"
)

type ctxKey string

const (
	ctxKeyLogger ctxKey = "logger"
	ctxKeyDB     ctxKey = "db"
	ctxKeyRedis  ctxKey = "redis"
	ctxKeyRender ctxKey = "render"
)

// GetLogger returns the logger associated with the context, or the global logger if none is set.
func GetLogger(ctx context.Context) *zap.Logger {
	if ctx == nil {
		ctx = context.Background()
	}
	if l, ok := ctx.Value(ctxKeyLogger).(*zap.Logger); ok {
		return l
	}
	return zap.L()
}

// WithLogger associates a logger with a context.
func WithLogger(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKeyLogger, l)
}

// WithLogger is a shorthand for associating a named sub-logger with a context.
func WithNamedLogger(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, ctxKeyLogger, GetLogger(ctx).Named(name))
}

// GetDB returns the DB associated with the context, or nil.
func GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(ctxKeyDB).(*gorm.DB)
	return db
}

// WithDB associates a DB with a context.
func WithDB(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxKeyDB, db)
}

// GetRedis returns the Redis client associated with the context, or nil.
func GetRedis(ctx context.Context) *redis.Client {
	r, _ := ctx.Value(ctxKeyRedis).(*redis.Client)
	return r
}

// WithRedis associates a Redis connection with a context.
func WithRedis(ctx context.Context, r *redis.Client) context.Context {
	return context.WithValue(ctx, ctxKeyRedis, r)
}

// GetRender returns the Render associated with the context, or nil.
func GetRender(ctx context.Context) *render.Render {
	rend, _ := ctx.Value(ctxKeyRender).(*render.Render)
	return rend
}

// WithRender associates a Render with a context.
func WithRender(ctx context.Context, rend *render.Render) context.Context {
	return context.WithValue(ctx, ctxKeyRender, rend)
}
