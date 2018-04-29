package jsonld

import (
	"encoding/json"
)

type Meta struct {
	Extra map[string]json.RawMessage `json:"-"`
}
