package api

import (
	"context"
	"net/http"
)

// Handler is a more convenient structure for an HTTP handler. By returning responses instead of
// using the ResponseWriter, we prevent errors arising from forgetting to return after a render
// call, and let content negotiation occur outside of the handlers themselves.
type Handler interface {
	HandleRequest(ctx context.Context, req *http.Request) Response
}

type funcHandler struct {
	h func(ctx context.Context, req *http.Request) Response
}

func (h funcHandler) HandleRequest(ctx context.Context, req *http.Request) Response {
	return h.h(ctx, req)
}

func HandlerFunc(h func(ctx context.Context, req *http.Request) Response) Handler {
	return &funcHandler{h: h}
}

// WrapHandler wraps a Handler into an http.Handler.
func WrapHandler(h Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		req = req.WithContext(ctx)
		RenderResponse(rw, req, h.HandleRequest(ctx, req))
	})
}
