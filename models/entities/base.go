package entities

import (
	"github.com/liclac/meow/jsonld"
)

type Base struct {
	jsonld.Meta

	ID   string   `json:"@id"`
	Type []string `json:"@type"`
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
