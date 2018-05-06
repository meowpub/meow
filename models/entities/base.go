package entities

import (
	"github.com/bwmarrin/snowflake"
	"github.com/liclac/meow/jsonld"
)

type Base struct {
	jsonld.Meta
	Snowflake snowflake.ID `json:"-"` // Only used internally.

	ID   string   `json:"@id"`
	Type []string `json:"@type"`
}

func (b *Base) SetSnowflake(id snowflake.ID) {
	b.Snowflake = id
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

func (b *Base) UnmarshalJSON(data []byte) error {
	return b.Meta.Unmarshal(data, b)
}

func (b Base) MarshalJSON() ([]byte, error) {
	return b.Marshal(b)
}
