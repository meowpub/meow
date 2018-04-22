package server

import (
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
