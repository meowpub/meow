package api

import (
	"net/http"

	"github.com/go-playground/form"
)

var formDecoder = form.NewDecoder()

func DecodeForm(req Request, v interface{}) error {
	if err := req.ParseForm(); err != nil {
		return Wrap(err, http.StatusBadRequest)
	}
	return Wrap(formDecoder.Decode(v, req.Form), http.StatusBadRequest)
}

func DecodeQuery(req Request, v interface{}) error {
	return Wrap(formDecoder.Decode(v, req.URL.Query()), http.StatusBadRequest)
}
