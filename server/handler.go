package server

import (
	"net/http"

	"github.com/liclac/meow/server/api"
	"go.uber.org/zap"
)

// WrapHandler wraps a Handler in an http.Handler.
func WrapHandler(h api.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		resp := h(ctx, req)

		RenderResponse(rw, req, resp)
	})
}

func RenderResponse(rw http.ResponseWriter, req *http.Request, resp api.Response) {
	ctx := req.Context()

	if resp.Error != nil {
		resp.Status = api.ErrorStatus(resp.Error)
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
}
