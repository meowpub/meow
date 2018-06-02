package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

var _ http.Handler = &Router{}

type Router struct {
	mux chi.Router
	mw  []Middleware
}

func NewRouter() *Router {
	return &Router{mux: chi.NewMux()}
}

func (r *Router) Use(mw ...Middleware) {
	r.mw = append(r.mw, mw...)
}

func (r *Router) wrap(h Handler) http.HandlerFunc {
	return WrapHandler(Chain(h, r.mw...))
}

func (r *Router) GET(path string, h Handler)        { r.mux.Get(path, r.wrap(h)) }
func (r *Router) HEAD(path string, h Handler)       { r.mux.Head(path, r.wrap(h)) }
func (r *Router) POST(path string, h Handler)       { r.mux.Post(path, r.wrap(h)) }
func (r *Router) PUT(path string, h Handler)        { r.mux.Put(path, r.wrap(h)) }
func (r *Router) DELETE(path string, h Handler)     { r.mux.Delete(path, r.wrap(h)) }
func (r *Router) ANY(path string, h Handler)        { r.mux.HandleFunc(path, r.wrap(h)) }
func (r *Router) Mount(path string, h http.Handler) { r.mux.Mount(path, h) }
func (r *Router) NotFound(h Handler)                { r.mux.NotFound(r.wrap(h)) }

func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(rw, req)
}
