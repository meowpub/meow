package api

import (
	"github.com/monzo/typhon"

	"github.com/meowpub/meow/ld"
)

// TODO: Do something with the Accept header here.
func EntityResponse(req typhon.Request, obj *ld.Object) typhon.Response {
	v, err := ld.CompactObject(req, obj)
	if err != nil {
		return typhon.Response{Error: err}
	}
	return req.Response(v)
}
