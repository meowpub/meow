package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	h := HandlerFunc(func(ctx context.Context, req *http.Request) Response {
		return Response{Data: "hi"}
	})

	req := httptest.NewRequest("GET", "/path", nil)
	ctx := req.Context()

	t.Run("No Middleware", func(t *testing.T) {
		assert.Equal(t, Response{Data: "hi"}, Chain(h).HandleRequest(ctx, req))
	})

	t.Run("One Middleware", func(t *testing.T) {
		mw := func(next Handler) Handler {
			return HandlerFunc(func(ctx context.Context, req *http.Request) Response {
				rsp := next.HandleRequest(ctx, req)
				rsp.Status = http.StatusNotFound
				return rsp
			})
		}
		assert.Equal(t, Response{Data: "hi", Status: 404}, Chain(h, mw).HandleRequest(ctx, req))
	})

	t.Run("Two Middleware", func(t *testing.T) {
		mw1 := func(next Handler) Handler {
			return HandlerFunc(func(ctx context.Context, req *http.Request) Response {
				rsp := next.HandleRequest(ctx, req)
				rsp.Status = http.StatusNotFound
				return rsp
			})
		}
		mw2 := func(next Handler) Handler {
			return HandlerFunc(func(ctx context.Context, req *http.Request) Response {
				rsp := next.HandleRequest(ctx, req)
				rsp.Data = "bye"
				return rsp
			})
		}
		assert.Equal(t, Response{Data: "bye", Status: 404}, Chain(h, mw1, mw2).HandleRequest(ctx, req))
	})
}
