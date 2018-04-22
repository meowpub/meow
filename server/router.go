package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/liclac/meow/config"
	"github.com/pkg/errors"
	"github.com/unrolled/render"
)

// New returns a new API router.
func New() http.Handler {
	r := chi.NewMux()
	r.Use(RenderMiddleware(render.Options{
		IndentJSON:    true,
		IsDevelopment: !config.IsProd(),
	}))
	r.Get("/", WrapHandler(Handler(HandleRoot)))
	return r
}

// HandleRoot serves the index page.
func HandleRoot(ctx context.Context, req *http.Request) Response {
	return Response{Error: errors.New("according to all known laws of aviation, there's no way a bee should be able to fly")}
}
