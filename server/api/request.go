package api

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/meowpub/meow/lib"
)

type Request struct{ *http.Request }

var _ context.Context = &Request{}

func (r Request) Deadline() (deadline time.Time, ok bool) {
	return r.Context().Deadline()
}

func (r Request) Done() <-chan struct{} {
	return r.Context().Done()
}

func (r Request) Err() error {
	return r.Context().Err()
}

func (r Request) Value(key interface{}) interface{} {
	return r.Context().Value(key)
}

func (r Request) WithContext(ctx context.Context) Request {
	return Request{r.Request.WithContext(unwrapContext(ctx))}
}

func (r Request) Context() context.Context {
	return unwrapContext(r.Request.Context())
}

func unwrapContext(ctx context.Context) context.Context {
	for {
		switch r := ctx.(type) {
		case Request:
			ctx = r.Request.Context()

		case *Request:
			ctx = r.Request.Context()

		default:
			return ctx
		}
	}
}

func (req Request) ResourceURL() url.URL {
	host, _, err := net.SplitHostPort(req.Host)
	if err != nil {
		host = req.Host
	}

	// Normalize the url
	url := *req.URL
	url.Scheme = "https"
	url.Host = host
	url = lib.NormalizeURL(url)

	return url
}

func (req Request) RootURL() url.URL {
	return lib.RootURL(req.ResourceURL())
}
