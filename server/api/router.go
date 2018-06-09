package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
	"github.com/liclac/meow/lib"
	"github.com/unrolled/render"
)

var _ http.Handler = &Router{}

type Router struct {
	rend *render.Render
	mux  chi.Router
	mw   []Middleware
}

func NewRouter(rendopt ...render.Options) *Router {
	return &Router{
		rend: render.New(rendopt...),
		mux:  chi.NewMux(),
	}
}

func (r *Router) Use(mw ...Middleware) {
	r.mw = append(r.mw, mw...)
}

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(rw, req)
}

func (r *Router) wrap(h Handler) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		r.Render(rw, req, Chain(h, r.mw...).HandleRequest(req.Context(), req))
	})
}

func (r *Router) Render(rw http.ResponseWriter, req *http.Request, resp Response) {
	L := lib.GetLogger(req.Context()).With(
		zap.String("method", req.Method),
		zap.String("path", req.URL.Path),
		zap.Int("status", resp.Status),
	)

	// Copy headers to the output first of all, regardless of what the response is.
	outHeader := rw.Header()
	for k, vs := range resp.Header() {
		outHeader[k] = vs
	}

	// If the response is an error, short-circuit and render that.
	if resp.Error != nil {
		L.Info("Request failed", zap.Error(resp.Error))
		r.RenderError(rw, req, resp.Status, resp.Error)
		return
	}

	// If no status is set, default to HTTP 200 OK.
	if resp.Status == 0 {
		resp.Status = http.StatusOK
	}

	// TODO: Do proper content negotiation here.
	var err error
	switch {
	case resp.Template != "":
		err = r.rend.HTML(rw, resp.Status, resp.Template, resp.Data)
	case resp.Data != nil:
		switch data := resp.Data.(type) {
		case string:
			err = r.rend.Text(rw, resp.Status, data)
		case []byte:
			err = r.rend.Data(rw, resp.Status, data)
		default:
			err = r.rend.JSON(rw, resp.Status, data)
		}
	default:
		rw.WriteHeader(resp.Status)
	}
	if err != nil {
		L.Error("Failed to render response", zap.Error(err))
		r.RenderError(rw, req, 0, err)
		return
	}

	L.Info("Request finished")
}

func (r *Router) RenderError(rw http.ResponseWriter, req *http.Request, status int, err error) {
	if status == 0 {
		status = ErrorStatus(err)
	}

	rw.WriteHeader(status)
	data, jsonerr := json.MarshalIndent(toErrorResponse(err), "", "  ")
	if jsonerr != nil {
		lib.GetLogger(req.Context()).DPanic("Failed to render response error",
			zap.String("method", req.Method),
			zap.String("path", req.URL.Path),
			zap.NamedError("jsonerr", jsonerr),
			zap.Error(err),
		)
		data = []byte(err.Error())
	}
	fmt.Fprintln(rw, string(data))
}

func (r *Router) GET(path string, h HandlerFunc)    { r.mux.Get(path, r.wrap(h)) }
func (r *Router) HEAD(path string, h HandlerFunc)   { r.mux.Head(path, r.wrap(h)) }
func (r *Router) POST(path string, h HandlerFunc)   { r.mux.Post(path, r.wrap(h)) }
func (r *Router) PUT(path string, h HandlerFunc)    { r.mux.Put(path, r.wrap(h)) }
func (r *Router) DELETE(path string, h HandlerFunc) { r.mux.Delete(path, r.wrap(h)) }
func (r *Router) ANY(path string, h HandlerFunc)    { r.mux.HandleFunc(path, r.wrap(h)) }
func (r *Router) Handle(path string, h Handler)     { r.mux.HandleFunc(path, r.wrap(h)) }
func (r *Router) Mount(path string, h http.Handler) { r.mux.Mount(path, h) }
func (r *Router) NotFound(h HandlerFunc)            { r.mux.NotFound(r.wrap(h)) }
