package api

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/meowpub/meow/models"
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
type statusError struct {
	Err        error
	StatusCode int
}

func Wrap(err error, code int) error {
	if err != nil {
		return statusError{err, code}
	}
	return nil
}

func (err statusError) Error() string { return err.Err.Error() }

// Cause satisfies causer.
func (err statusError) Cause() error { return err.Err }

// ErrorStatus returns the status code for an error (defaults to 500, 200 for a nil error).
func ErrorStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if e, ok := err.(statusError); ok {
		return e.StatusCode
	}

	if models.IsNotFound(err) {
		return http.StatusNotFound
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
