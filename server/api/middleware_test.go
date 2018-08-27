package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain(t *testing.T) {
	h := HandlerFunc(func(req Request) Response {
		return Response{Data: "hi"}
	})

	req := Request{httptest.NewRequest("GET", "/path", nil)}

	t.Run("No Middleware", func(t *testing.T) {
		assert.Equal(t, Response{Data: "hi"}, Chain(h).HandleRequest(req))
	})

	t.Run("One Middleware", func(t *testing.T) {
		mw := func(next Handler) Handler {
			return HandlerFunc(func(req Request) Response {
				rsp := next.HandleRequest(req)
				rsp.Status = http.StatusNotFound
				return rsp
			})
		}
		assert.Equal(t, Response{Data: "hi", Status: 404}, Chain(h, mw).HandleRequest(req))
	})

	t.Run("Two Middleware", func(t *testing.T) {
		mw1 := func(next Handler) Handler {
			return HandlerFunc(func(req Request) Response {
				rsp := next.HandleRequest(req)
				rsp.Status = http.StatusNotFound
				return rsp
			})
		}
		mw2 := func(next Handler) Handler {
			return HandlerFunc(func(req Request) Response {
				rsp := next.HandleRequest(req)
				rsp.Data = "bye"
				return rsp
			})
		}
		assert.Equal(t, Response{Data: "bye", Status: 404}, Chain(h, mw1, mw2).HandleRequest(req))
	})
}
