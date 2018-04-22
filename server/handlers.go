package server

import (
	"context"
	"net/http"

	"go.uber.org/zap"
)

// Handler is a more convenient structure for an HTTP handler. By returning responses instead of
// using the ResponseWriter, we prevent errors arising from forgetting to return after a render
// call, and let content negotiation occur outside of the handlers themselves.
type Handler func(ctx context.Context, req *http.Request) Response

// WrapHandler wraps a Handler in an http.Handler.
func WrapHandler(h Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		resp := h(ctx, req)
		if resp.Error != nil {

			resp.Status = ErrorStatus(resp.Error)
			resp.Data = toErrorResponse(resp.Error)
		}
		if resp.Data == nil {
			return
		}

		var err error
		switch data := resp.Data.(type) {
		case string:
			err = GetRender(ctx).Text(rw, resp.Status, data)
		case []byte:
			err = GetRender(ctx).Data(rw, resp.Status, data)
		default:
			err = GetRender(ctx).JSON(rw, resp.Status, data)
		}

		if err != nil {
			rw.Write([]byte(err.Error()))
			zap.L().Error("failed to render response", zap.Error(err))
		}
	})
}
