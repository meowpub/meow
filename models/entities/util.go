package entities

import (
	"context"
	"net/http"

	"github.com/pkg/errors"

	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/server/api"
)

func as2(postfix string) string {
	return "https://www.w3.org/ns/activitystreams#" + postfix
}

func compact(ctx context.Context, data interface{}) (interface{}, error) {
	return jsonld.Compact(lib.GetHttpClient(ctx), data.(map[string]interface{}), "", "https://www.w3.org/ns/activitystreams")
}

func handleEntityGetRequest(ctx context.Context, e Entity, req *http.Request) api.Response {
	if req.Method != http.MethodGet && req.Method != http.MethodOptions && req.Method != http.MethodHead {
		return api.Response{
			Status: http.StatusMethodNotAllowed,
			Error:  errors.New("Unsupported request method " + req.Method),
		}
	}

	d, err := e.Hydrate(ctx)
	if err != nil {
		return api.ErrorResponse(err)
	}

	o, err := compact(ctx, d)
	if err != nil {
		return api.ErrorResponse(err)
	}

	return api.Response{
		Data: o,
	}
}
