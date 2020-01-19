package api

import (
	"net"
	"net/http"

	"github.com/monzo/typhon"
	"github.com/pkg/errors"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
)

var Service = typhon.Service(Handler)

func Handler(req typhon.Request) typhon.Response {
	u := *req.URL
	u.Host, _, _ = net.SplitHostPort(req.Host)
	if u.Host == "" {
		u.Host = req.Host
	}
	entity, err := Lookup(req, u)
	if err != nil {
		return typhon.Response{Error: err}
	}
	if entity == nil {
		return typhon.Response{Error: lib.Error(http.StatusNotFound, "Not Found")}
	}
	switch req.Method {
	case http.MethodGet:
		return entity.GET(req)
	default:
		return typhon.Response{Error: lib.Error(http.StatusMethodNotAllowed, "Method Not Allowed")}
	}
}

// Returns the HTTP status code for an error.
func ErrorCode(err error) int {
	if code := errorCode(err); code != 0 {
		return code
	}
	return http.StatusInternalServerError
}

func errorCode(err error) int {
	if err, ok := err.(lib.Err); ok {
		return err.StatusCode
	}
	if models.IsNotFound(err) {
		return http.StatusNotFound
	}
	if cause := errors.Cause(err); cause != err {
		return errorCode(cause)
	}
	return 0
}
