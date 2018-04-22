package server

import (
	"context"
	"net/http"

	"github.com/unrolled/render"
)

const ctxKeyRender ctxKey = "render"

// RenderMiddleware adds a Render to request contexts.
func RenderMiddleware(opts ...render.Options) func(next http.Handler) http.Handler {
	r := render.New(opts...)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			req = req.WithContext(WithRender(req.Context(), r))
			next.ServeHTTP(rw, req)
		})
	}
}

// WithRender associates a Render with a context.
func WithRender(ctx context.Context, r *render.Render) context.Context {
	return context.WithValue(ctx, ctxKeyRender, r)
}

// GetRender returns the Render associated with a context.
func GetRender(ctx context.Context) *render.Render {
	r, _ := ctx.Value(ctxKeyRender).(*render.Render)
	return r
}
