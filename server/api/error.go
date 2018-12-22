package api

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
)

// Lifted from pkg/errors.
type causer interface {
	Cause() error
}

func cause(err error) error {
	if cerr, ok := err.(causer); ok {
		return cerr.Cause()
	}
	return nil
}

// Lifted from pkg/errors.
type stackTracer interface {
	StackTrace() errors.StackTrace
}

func stackTrace(err error) errors.StackTrace {
	if serr, ok := err.(stackTracer); ok {
		return serr.StackTrace()
	}
	if cerr := cause(err); cerr != nil {
		return stackTrace(cerr)
	}
	return nil
}

// ErrorStatus returns the status code for an error (defaults to 500, 200 for a nil error).
func ErrorStatus(err error) int {
	if status := errorStatus(err); status != 0 {
		return status
	}
	return http.StatusInternalServerError
}

func errorStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if e, ok := err.(lib.Err); ok {
		return e.StatusCode
	}

	if models.IsNotFound(err) {
		return http.StatusNotFound
	}

	if e, ok := err.(causer); ok {
		return ErrorStatus(e.Cause())
	}

	return 0
}
