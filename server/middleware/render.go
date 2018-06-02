package middleware

import (
	"context"
	"net/http"

	"github.com/unrolled/render"

	"github.com/liclac/meow/lib"
	"github.com/liclac/meow/server/api"
)

func AddRender(rend *render.Render) func(next api.Handler) api.Handler {
	return func(next api.Handler) api.Handler {
		return api.HandlerFunc(func(ctx context.Context, req *http.Request) api.Response {
			ctx = lib.WithRender(ctx, rend)
			req = req.WithContext(ctx)
			return next.HandleRequest(ctx, req)
		})
	}
}
