package middleware

import (
	"context"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models"
	"github.com/liclac/meow/models/entities"
	"github.com/liclac/meow/server/api"
)

func AddDB(db *gorm.DB) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(ctx context.Context, req *http.Request) api.Response {
			ctx = lib.WithDB(ctx, db)
			req = req.WithContext(ctx)
			return next.HandleRequest(ctx, req)
		})
	}
}

func AddRedis(r *redis.Client) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(ctx context.Context, req *http.Request) api.Response {
			ctx = lib.WithRedis(ctx, r)
			req = req.WithContext(ctx)
			return next.HandleRequest(ctx, req)
		})
	}
}

func AddStores() func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(ctx context.Context, req *http.Request) api.Response {
			stores := models.NewStores(
				lib.GetDB(ctx),
				lib.GetRedis(ctx),
				config.RedisKeyspace(),
			)
			estore := entities.NewStore(stores.Entities())

			ctx = models.WithStores(ctx, stores)
			ctx = entities.WithStore(ctx, estore)
			req = req.WithContext(ctx)
			return next.HandleRequest(ctx, req)
		})
	}
}
