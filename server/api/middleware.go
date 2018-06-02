package api

type Middleware func(next Handler) Handler

func Chain(h Handler, mw ...Middleware) Handler {
	if len(mw) == 0 {
		return h
	}
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return h
}
