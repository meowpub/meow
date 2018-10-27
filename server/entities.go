package server

import (
	"context"
	"path"

	"go.uber.org/zap"

	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
	"github.com/meowpub/meow/server/api"
)

var _ api.Traversible = EntityNode{}

// Wraps an Entity in a Traversible node, which can serve itself to a HTTP response.
type EntityNode struct {
	Entity *models.Entity
}

// Uses Lookup to traverse to a child node.
// TODO: Implement support for "magic nodes" generated at runtime, eg. inboxes.
func (n EntityNode) Traverse(ctx context.Context, pathElement string) (api.Handler, error) {
	return Lookup(ctx, path.Join(n.Entity.Obj.ID(), pathElement))
}

// Renders the entity to an HTTP response.
// TODO: Re-implement Hydration from the old codebase, actually instantiate concrete,
// typed entities (eg. ld/ns/as.Note) to allow for type-specific behaviour.
func (n EntityNode) HandleRequest(req api.Request) api.Response {
	return api.Response{Data: n.Entity.Obj}
}

// Called by the Router to return a node at an arbitrary point in the tree.
// This is called both by EntityNode.Traverse and the router.
func Lookup(ctx context.Context, url string) (api.Traversible, error) {
	e, err := models.GetStores(ctx).Entities().GetByID(url)
	if err != nil {
		lib.GetLogger(ctx).Debug("Lookup failed", zap.String("url", url), zap.Error(err))
		return nil, err
	}
	if e == nil {
		lib.GetLogger(ctx).Debug("Lookup found nothing", zap.String("url", url))
		return nil, nil
	}
	return EntityNode{e}, nil
}
