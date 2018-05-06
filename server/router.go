package server

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/models/entities"
	"github.com/liclac/meow/server/api"
	"github.com/unrolled/render"
)

// New returns a new API router.
func New() http.Handler {
	r := chi.NewMux()
	r.Use(DBMiddleware())
	r.Use(RenderMiddleware(render.Options{
		IndentJSON:    true,
		IsDevelopment: !config.IsProd(),
	}))

	r.Get("/", WrapHandler(RouteRequest))
	r.NotFound(WrapHandler(RouteRequest))

	return r
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
