package api

// Response is a structured response.
type Response struct {
	Status   int         // HTTP status code.
	Error    error       // Error, if any.
	Template string      // Template for HTML responses.
	Data     interface{} // Data, marshalled (JSON) or given to a template (HTML).
}

func ErrorResponse(err error) Response {
	return Response{
		Error: err,
	}
}
