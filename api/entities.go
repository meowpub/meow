package api

import (
	"context"
	"net/url"
	"strings"

	"github.com/monzo/typhon"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
)

type Node interface {
	GET(typhon.Request) typhon.Response
}

type Object struct{ *ld.Object }

func (o Object) GET(req typhon.Request) typhon.Response {
	return EntityResponse(req, o.Object)
}

func Instantiate(obj *ld.Object) Node {
	return Object{obj}
}

func Lookup(ctx context.Context, reqURL url.URL) (Node, error) {
	// Normalise the URL, forcing it to https:// and removing any authentication or queries.
	reqURL = url.URL{Scheme: "https", Host: reqURL.Host, Path: reqURL.Path}
	lib.GetLogger(ctx).Debug("Looking up entity...", zap.String("id", reqURL.String()))

	entities := models.GetStores(ctx).Entities()
	e, err := entities.GetByID(reqURL.String())
	if err != nil {
		if models.IsNotFound(err) && strings.HasSuffix(reqURL.Path, "/") {
			newURL := reqURL
			newURL.Path = strings.TrimSuffix(newURL.Path, "/")
			return Lookup(ctx, newURL)
		}
		return nil, errors.Wrap(err, reqURL.String())
	}
	return Instantiate(e.Obj), nil
}
