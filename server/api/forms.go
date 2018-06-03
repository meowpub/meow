package api

import (
	"net/http"

	"github.com/go-playground/form"
)

var formDecoder = form.NewDecoder()

func DecodeForm(req *http.Request, v interface{}) error {
	if err := req.ParseForm(); err != nil {
		return Wrap(err, http.StatusBadRequest)
	}
	return formDecoder.Decode(v, req.Form)
}
