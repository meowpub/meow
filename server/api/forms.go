package api

import (
	"net/http"

	"github.com/meowpub/meow/lib"

	"github.com/go-playground/form"
)

var formDecoder = form.NewDecoder()

func DecodeForm(req Request, v interface{}) error {
	if err := req.ParseForm(); err != nil {
		return lib.Code(err, http.StatusBadRequest)
	}
	return lib.Code(formDecoder.Decode(v, req.Form), http.StatusBadRequest)
}

func DecodeQuery(req Request, v interface{}) error {
	return lib.Code(formDecoder.Decode(v, req.URL.Query()), http.StatusBadRequest)
}
