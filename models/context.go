package models

import (
	"context"
)

type ctxKey string

const ctxKeyStores ctxKey = "stores"

// GetStores returns the Stores object associated with the context.
func GetStores(ctx context.Context) Stores {
	s, _ := ctx.Value(ctxKeyStores).(Stores)
	return s
}

// WithStores associates a Stores object with a context.
func WithStores(ctx context.Context, stores Stores) context.Context {
	return context.WithValue(ctx, ctxKeyStores, stores)
}
