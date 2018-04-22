package server

import (
	"net/http"
)

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
