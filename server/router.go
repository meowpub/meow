package server

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/unrolled/render"

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models/entities"
	"github.com/liclac/meow/server/api"
	"github.com/liclac/meow/server/middleware"
)

// New returns a new API router.
func New(db *gorm.DB, r *redis.Client) http.Handler {
	mux := chi.NewMux()
	mux.Use(middleware.AddDB(db))
	mux.Use(middleware.AddRedis(r))
	mux.Use(middleware.AddStores())
	mux.Use(middleware.AddRender(render.New(render.Options{
		IndentJSON:    true,
		IsDevelopment: !config.IsProd(),
	})))

	mux.Get("/", WrapHandler(RouteRequest))
	mux.NotFound(WrapHandler(RouteRequest))

	return mux
}

func RouteRequest(ctx context.Context, req *http.Request) api.Response {
	host, _, err := net.SplitHostPort(req.Host)
	if err != nil {
		return api.ErrorResponse(err)
	}

	url := *req.URL
	url.Scheme = "https"
	url.Host = host
	url = lib.NormalizeURL(url)

	store := entities.GetStore(ctx)
	ent, err := store.GetByID(url.String())

	if err != nil {
		return api.ErrorResponse(err)
	} else {
		return api.Response{
			Data: ent,
		}
	}
}
