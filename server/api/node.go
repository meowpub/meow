package api

import (
	"context"
	"net/http"
)

var _ Traversible = Node{}

// Nodes are fixed Traversibles, useful for quickly defining a hardcoded path.
type Node struct {
	Self     Handler
	Children map[string]Handler
}

func (n Node) HandleRequest(ctx context.Context, req *http.Request) Response {
	return n.Self.HandleRequest(ctx, req)
}

func (n Node) Traverse(ctx context.Context, pathElement string) (Handler, error) {
	if h, ok := n.Children[pathElement]; ok {
		return h, nil
	}
	if t, ok := n.Self.(Traversible); ok {
		return t.Traverse(ctx, pathElement)
	}
	return nil, nil
}
