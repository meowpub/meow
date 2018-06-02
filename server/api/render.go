package api

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/liclac/meow/lib"
)

func RenderError(rw http.ResponseWriter, req *http.Request, status int, rspErr error) {
	ctx := req.Context()
	rsp := toErrorResponse(rspErr)

	// Default to the status pulled from the error, which defaults to 500.
	if status == 0 {
		status = ErrorStatus(rspErr)
	}

	l := lib.GetLogger(ctx).WithOptions(zap.Fields(
		zap.String("path", req.URL.Path),
		zap.String("method", req.Method),
		zap.Int("status", status),
		zap.Error(rspErr),
	))
	l.Debug("rendering error")

	data, err := json.MarshalIndent(rsp, "", "  ")
	if err != nil {
		l.DPanic("COULDN'T MARSHAL ERROR", zap.NamedError("jsonerr", err))
		data = []byte(rspErr.Error())
	}

	rw.WriteHeader(status)
	rw.Write(data)
}

func RenderResponse(rw http.ResponseWriter, req *http.Request, resp Response) {
	ctx := req.Context()
	l := lib.GetLogger(ctx).WithOptions(zap.Fields(
		zap.String("path", req.URL.Path),
		zap.String("method", req.Method),
	))

	// If the response is an error, render just the status.
	if resp.Error != nil {
		l.Debug("response is an error", zap.Error(resp.Error))
		RenderError(rw, req, resp.Status, resp.Error)
		return
	}

	// Default to a HTTP 200 OK response if no status is set.
	if resp.Status == 0 {
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

	// Render an error if rendering the contents failed for whatever reason.
	if rerr != nil {
		RenderError(rw, req, http.StatusInternalServerError, rerr)
	}
}
