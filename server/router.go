package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/liclac/meow/config"
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
	url := *req.URL
	url.RawQuery = ""
	url.ForceQuery = false
	url.Fragment = ""

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
