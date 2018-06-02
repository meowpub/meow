package api

import (
	"context"
)

const ctxKeyTraversal ctxKey = "traversal"

// TraversalContext is the context traversed so far in the processing of
// this request
type TraversalContext struct {
	// Path up until the found handler
	TraversedPath []Traversible
	FoundHandler  Handler
	RemainingPath []string
}

// A traversible object is one from which children can be 'walked'
type Traversible interface {
	Handler
	Traverse(ctx context.Context, pathElement string) (Handler, error)
}

// TraverseFrom traverses a path from a root object
func TraverseFrom(ctx context.Context, root Handler, path []string) (*TraversalContext, error) {
	tc := &TraversalContext{
		TraversedPath: []Traversible{},
		FoundHandler:  root,
		RemainingPath: path,
	}

	for len(tc.RemainingPath) > 0 {
		if tc.RemainingPath[0] == "" {
			tc.RemainingPath = tc.RemainingPath[1:]
			continue
		}

		cur, ok := tc.FoundHandler.(Traversible)
		if !ok {
			break
		}

		next, err := cur.Traverse(ctx, tc.RemainingPath[0])
		if err != nil {
			return nil, err
		}

		if next == nil {
			break
		}

		tc.TraversedPath = append(tc.TraversedPath, cur)
		tc.FoundHandler = next
		tc.RemainingPath = tc.RemainingPath[1:]
	}

	return tc, nil
}

func WithTraversalContext(ctx context.Context, cx *TraversalContext) context.Context {
	return context.WithValue(ctx, ctxKeyTraversal, cx)
}

func GetTraversalContext(ctx context.Context) *TraversalContext {
	return ctx.Value(ctxKeyTraversal).(*TraversalContext)
}
