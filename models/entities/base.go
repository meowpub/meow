package entities

import (
	"context"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/jsonld"
	"github.com/meowpub/meow/server/api"
)

type Base struct {
	jsonld.Meta
	Snowflake snowflake.ID `json:"-"` // Only used internally.

	ID   string   `json:"@id"`
	Type []string `json:"@type"`
}

func (b Base) Init(_ context.Context) error {
	return nil
}

func (b Base) GetSnowflake() snowflake.ID {
	return b.Snowflake
}

func (b Base) GetID() string {
	return b.ID
}

func (b Base) GetType() []string {
	return b.Type
}

// api.Traversible
func (b *Base) Traverse(ctx context.Context, pathElement string) (api.Handler, error) {
	store := GetStore(ctx)

	id := b.GetID() + "/" + pathElement

	return store.GetByID(id)
}
