package lib

import (
	"context"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Error attaches a status code to an error.
type Err struct {
	Err        error
	StatusCode int
}

// Wraps an error with an error code.
func Code(err error, code int) error {
	if err != nil {
		return Err{err, code}
	}
	return nil
}

// Wraps an error with an error code and message.
func Wrap(err error, code int, s string) error {
	return errors.Wrap(Code(err, code), s)
}

// Wraps an error with an error code and message.
func Wrapf(err error, code int, s string, args ...interface{}) error {
	return errors.Wrapf(Code(err, code), s, args...)
}

// Returns an error with the given message and status code.
func Error(code int, s string) error {
	return Code(errors.New(s), code)
}

// Returns an error with the given message format and status code.
func Errorf(code int, s string, args ...interface{}) error {
	return Code(errors.Errorf(s, args...), code)
}

// Satisfies error.
func (err Err) Error() string { return err.Err.Error() }

// Satisfies errors.causer.
func (err Err) Cause() error { return err.Err }

// Report logs an error to the global logger. Useful for `defer conn.Close()`-type constructs,
// where there's not really anything useful to do with the error, but you still want to log it.
func Report(ctx context.Context, err error, msg string) {
	if err != nil {
		GetLogger(ctx).Error(msg, zap.Error(err))
	}
}
