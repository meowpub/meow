package api

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
)

// Handler is a more convenient structure for an HTTP handler. By returning responses instead of
// using the ResponseWriter, we prevent errors arising from forgetting to return after a render
// call, and let content negotiation occur outside of the handlers themselves.
type Handler interface {
	HandleRequest(req Request) Response
}

// An Unwrappable handler can be unwrapped to find a nested handler
// (In general HandleRequest should delegate straight to the wrapped handler)
type UnwrappableHandler interface {
	Handler
	UnwrapHandler() Handler
}

// HandlerFunc lets a function act as a Handler.
type HandlerFunc func(Request) Response

func (h HandlerFunc) HandleRequest(req Request) Response {
	return h(req)
}

// Implement MarshalJSON for HandlerFunc to generate useful logging
func (h HandlerFunc) MarshalJSON() ([]byte, error) {
	name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
	return json.Marshal(fmt.Sprintf("<function %s>", name))
}

// UnwrapHandler repeatedly trys to unwrap a handler
func UnwrapHandler(h Handler) Handler {
	for {
		if n, ok := h.(UnwrappableHandler); ok {
			if nh := n.UnwrapHandler(); nh != nil {
				h = nh
			} else {
				return h
			}
		} else {
			return h
		}
	}
}
