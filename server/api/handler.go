package api

// Handler is a more convenient structure for an HTTP handler. By returning responses instead of
// using the ResponseWriter, we prevent errors arising from forgetting to return after a render
// call, and let content negotiation occur outside of the handlers themselves.
type Handler interface {
	HandleRequest(req Request) Response
}

// HandlerFunc lets a function act as a Handler.
type HandlerFunc func(Request) Response

func (h HandlerFunc) HandleRequest(req Request) Response {
	return h(req)
}
