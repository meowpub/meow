package middleware

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
)

func AddDB(db *gorm.DB) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(req api.Request) api.Response {
			ctx := lib.WithDB(req.Context(), db)
			req = req.WithContext(ctx)
			return next.HandleRequest(req.WithContext(lib.WithDB(req.Context(), db)))
		})
	}
}

func AddRedis(r *redis.Client) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(req api.Request) api.Response {
			ctx := lib.WithRedis(req.Context(), r)
			req = req.WithContext(ctx)
			return next.HandleRequest(req)
		})
	}
}

func AddStores() func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(req api.Request) api.Response {
			stores := models.NewStores(
				lib.GetDB(req),
				lib.GetRedis(req),
				config.RedisKeyspace(),
			)

			ctx := models.WithStores(req.Context(), stores)
			req = req.WithContext(ctx)
			return next.HandleRequest(req)
		})
	}
}
