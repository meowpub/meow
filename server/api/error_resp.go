package api

import (
	"fmt"
	"runtime"
)

type stackFrame struct {
	Func string `json:"func"`
	File string `json:"file"`
}

// errorResponse is the Data type given to Error responses.
type errorResponse struct {
	StatusCode int          `json:"status_code"`
	Error      string       `json:"error"`
	StackTrace []stackFrame `json:"stack_trace,omitempty"`
}

func toErrorResponse(err error) errorResponse {
	var stack []stackFrame
	for _, frame := range StackTrace(err) {
		fn := runtime.FuncForPC(uintptr(frame))
		file, line := fn.FileLine(uintptr(frame))
		stack = append(stack, stackFrame{
			Func: fn.Name(),
			File: fmt.Sprintf("%s:%d", file, line),
		})
	}
	return errorResponse{
		StatusCode: ErrorStatus(err),
		Error:      err.Error(),
		StackTrace: stack,
	}
}
