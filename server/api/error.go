package api

import (
	"net/http"

	"github.com/pkg/errors"
)

// From pkg/errors; unexported but stable.
type causer interface {
	Cause() error
}

// From pkg/errors; unexported but stable.
type stackTracer interface {
	StackTrace() errors.StackTrace
}

// Error attaches a status code to an error.
type Error struct {
	error
	StatusCode int
}

// Cause satisfies causer.
func (err Error) Cause() error { return err.error }

// ErrorStatus returns the status code for an error (defaults to 500).
func ErrorStatus(err error) int {
	if e, ok := err.(Error); ok {
		return e.StatusCode
	}
	if e, ok := err.(causer); ok {
		return ErrorStatus(e.Cause())
	}
	return http.StatusInternalServerError
}

// StackTrace returns the stack trace for an error or any parent errors, if there is one.
func StackTrace(err error) errors.StackTrace {
	if e, ok := err.(stackTracer); ok {
		return e.StackTrace()
	}
	if e, ok := err.(causer); ok {
		return StackTrace(e.Cause())
	}
	return nil
}
