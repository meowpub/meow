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

// HandlerFunc lets a function act as a Handler.
type HandlerFunc func(ctx context.Context, req *http.Request) Response

func (h HandlerFunc) HandleRequest(ctx context.Context, req *http.Request) Response {
	return h(ctx, req)
}
