package api

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/liclac/meow/lib"
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
		RenderResponse(rw, req, h.HandleRequest(req.Context(), req))
	})
}

func RenderResponse(rw http.ResponseWriter, req *http.Request, resp Response) {
	ctx := req.Context()
	l := lib.GetLogger(ctx).WithOptions(zap.Fields(
		zap.String("path", req.URL.Path),
		zap.String("method", req.Method),
	))

	if resp.Error != nil {
		l.Debug("response is an error", zap.Error(resp.Error))
		if resp.Status == 0 {
			resp.Status = ErrorStatus(resp.Error)
			l.Debug("status code set from error", zap.Int("status", resp.Status))
		}
		resp.Data = toErrorResponse(resp.Error)
	}

	if resp.Status == 0 {
		l.Debug("status = 0, setting status = 200")
		resp.Status = 200
	}

	var rerr error
	switch {
	case resp.Template != "":
		l.Debug("rendering template", zap.String("template", resp.Template))
		rerr = lib.GetRender(ctx).HTML(rw, resp.Status, resp.Template, resp.Data)
	case resp.Data != nil:
		switch data := resp.Data.(type) {
		case string:
			l.Debug("rendering verbatim text")
			rerr = lib.GetRender(ctx).Text(rw, resp.Status, data)
		case []byte:
			l.Debug("rendering verbatim data")
			rerr = lib.GetRender(ctx).Data(rw, resp.Status, data)
		default:
			l.Debug("rendering json")
			rerr = lib.GetRender(ctx).JSON(rw, resp.Status, data)
		}
	default:
		l.Debug("rendering empty response")
		rerr = lib.GetRender(ctx).Data(rw, resp.Status, nil)
	}

	if rerr != nil {
		l.Error("FAILED TO RENDER RESPONSE", zap.Error(rerr))
		data, _ := json.MarshalIndent(toErrorResponse(rerr), "", "  ")
		rw.Write([]byte(data))
	}
}
