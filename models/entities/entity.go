package entities

import (
	"context"
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/meowpub/meow/server/api"
	"github.com/pkg/errors"
)

var (
	kinds map[string]*EntityKind = make(map[string]*EntityKind)
)

type Entity interface {
	api.Traversible

	// Initialize a new entity
	Init(context.Context) error

	// Hydrates the object
	Hydrate(context.Context, []snowflake.ID) (interface{}, error)

	// GetSnowflake returns the internal snowflake of the entity, eg. 353894652568535040.
	GetSnowflake() snowflake.ID

	// GetID returns the public ID of the entity, eg. "https://example.com/@jsmith".
	GetID() string

	// GetType returns the type of the entity, eg. ["https://www.w3.org/ns/activitystreams#Person"].
	// An entity may have multiple types, in which case it will behave as all of them at once.
	GetType() []string

	// GetKind returns the Kind of an entity
	GetKind() *EntityKind
}

// EntityKind is the 'meta type' of an Entity
type EntityKind struct {
	// The name of the kind (as stored in the database)
	Name string

	// Make a new entity of this kind
	// This should only do basic initialization - any complex
	// initialization should be done in Entity.Init
	New func(flake snowflake.ID, id string) (Entity, error)

	// Unmarshall this object into an Entity
	Unmarshall func(snowflake.ID, map[string]interface{}) (Entity, error)

	// Marshall this object into an Entity
	Marshall func(Entity) (map[string]interface{}, error)
}

func RegisterKind(kind *EntityKind) {
	if _, ok := kinds[kind.Name]; ok {
		panic(fmt.Sprintf("Attempt to register entity kind '%s' twice", kind.Name))
	}

	kinds[kind.Name] = kind
}

func GetKind(name string) (*EntityKind, error) {
	if kind, ok := kinds[name]; ok {
		return kind, nil
	} else {
		return nil, errors.Errorf("Unable to find entity kind '%s'", name)
	}
}
