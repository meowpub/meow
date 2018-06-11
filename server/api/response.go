package api

import (
	"bytes"
	"context"
	"net/http"

	"github.com/pkg/errors"
)

var _ http.ResponseWriter = &Response{}

// Hydratable indicates that this object can be "hydrated"
// before rendering
type Hydratable interface {
	// Hydrates the object
	Hydrate(context.Context) (interface{}, error)
}

// Response is a structured response. Implements http.ResponseWriter.
type Response struct {
	Status   int         // HTTP status code.
	Error    error       // Error, if any.
	Template string      // Template for HTML responses.
	Data     interface{} // Data, marshalled (JSON) or given to a template (HTML).

	header http.Header // Header values.
}

func (rsp *Response) Header() http.Header {
	if rsp.header == nil {
		rsp.header = make(http.Header)
	}
	return rsp.header
}

func (rsp *Response) WriteHeader(status int) {
	rsp.Status = status
}

func (rsp *Response) Write(b []byte) (int, error) {
	// If Data is set, append to it if it's []byte, else error out.
	var raw []byte
	if rsp.Data != nil {
		bytes, ok := rsp.Data.([]byte)
		if !ok {
			return 0, errors.Errorf("attempted to write to a non-binary body")
		}
		raw = bytes
	}

	// Make a buffer and write to it. This can definitely be optimised.
	buf := bytes.NewBuffer(raw)
	n, err := buf.Write(b)
	if err != nil {
		return n, err
	}
	rsp.Data = buf.Bytes()

	return n, nil
}

func RedirectResponse(url string) Response {
	resp := Response{Status: http.StatusTemporaryRedirect}
	resp.Header().Set("Location", url)
	return resp
}

func ErrorResponse(err error) Response {
	return Response{
		Error: err,
	}
}
