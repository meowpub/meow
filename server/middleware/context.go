package middleware

import (
	"net/http"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models"
	"github.com/liclac/meow/models/entities"
	"github.com/unrolled/render"
)

func AddDB(db *gorm.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(rw, req.WithContext(lib.WithDB(req.Context(), db)))
		})
	}
}

func AddRedis(r *redis.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(rw, req.WithContext(lib.WithRedis(req.Context(), r)))
		})
	}
}

func AddRender(rend *render.Render) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			next.ServeHTTP(rw, req.WithContext(lib.WithRender(req.Context(), rend)))
		})
	}
}

func AddStores() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			stores := models.NewStores(lib.GetDB(ctx), lib.GetRedis(ctx), config.RedisKeyspace())
			next.ServeHTTP(rw, req.WithContext(models.WithStores(ctx, stores)))
		})
	}
}

func AddEntityStore() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			stores := models.GetStores(ctx)
			store := entities.NewStore(stores.Entities())
			next.ServeHTTP(rw, req.WithContext(entities.WithStore(ctx, store)))
		})
	}
}
