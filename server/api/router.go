package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/unrolled/render"
	"go.uber.org/zap"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/ld/ns"
	"github.com/meowpub/meow/lib"
)

var _ Handler = &Router{}
var _ http.Handler = &Router{}

type LookupFunc func(ctx context.Context, normalizedURL *url.URL) (Traversible, error)

type Router struct {
	// Returns the entity at the given path. The given URL will always be normalized.
	lookup LookupFunc

	// Renderer used to serve responses.
	rend *render.Render

	// Hard routes that will override any attempt to traverse. Usually Node objects.
	// Note that these will only be looked at if the root object exists.
	hardRoutes map[string]Handler

	// Middleware stack.
	mw []Middleware
}

func NewRouter(lookup LookupFunc, opts ...render.Options) *Router {
	renderOpts := render.Options{}
	if len(opts) > 0 {
		renderOpts = opts[0]
	}
	renderOpts.DisableHTTPErrorRendering = true

	return &Router{
		lookup:     lookup,
		rend:       render.New(renderOpts),
		hardRoutes: make(map[string]Handler),
	}
}

func (r *Router) Mount(path string, h Handler) {
	r.hardRoutes[path] = h
}

func (r *Router) Use(mw ...Middleware) {
	r.mw = append(r.mw, mw...)
}

func (r Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	apiReq := Request{req}
	r.Render(rw, apiReq, r.HandleRequest(apiReq))
}

func (r Router) HandleRequest(req Request) Response {
	// Just wrap handleRequest in our own middleware stack and invoke that.
	// This ensures that lookups, etc. get middlewares applied properly.
	return Chain(HandlerFunc(r.handleRequest), r.mw...).HandleRequest(req)
}

func (r Router) handleRequest(req Request) Response {
	url := req.ResourceURL()
	rootUrl := lib.RootURL(url)

	// Find the root for the domain, make a Node out of it to reuse existing hard route logic.
	root, err := r.lookup(req, &rootUrl)
	if err != nil {
		return Response{Error: err}
	}
	if root == nil {
		return Response{Status: http.StatusNotFound}
	}
	node := Node{root, r.hardRoutes}

	// Traverse the node, find a handler.
	path := strings.Split(url.Path, "/")
	tc, err := TraverseFrom(req, node, path)
	if err != nil {
		return Response{Error: err}
	}

	ctx := WithTraversalContext(req.Context(), tc)
	req = req.WithContext(ctx)
	return tc.FoundHandler.HandleRequest(req)
}

func (r *Router) Render(rw http.ResponseWriter, req Request, resp Response) {
	L := lib.GetLogger(req).With(
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
		case ld.Entity:
			err = r.renderJSONLD(rw, resp.Status, data)
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

func (r *Router) RenderError(rw http.ResponseWriter, req Request, status int, err error) {
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

func (r *Router) renderJSONLD(rw http.ResponseWriter, status int, data ld.Entity) error {
	compact, err := ld.Compact(&http.Client{}, data.Obj().V, "", ns.AS.ID)
	if err != nil {
		return err
	}
	return r.rend.JSON(rw, status, compact)
}
