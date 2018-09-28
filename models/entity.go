package models

import (
	"encoding/json"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/lib"
)

//go:generate mockgen -package=models -source=entity.go -destination=entity.mock.go

var entityOnConflict = genOnConflict(Entity{}, "id")

// An Entity represents a raw, database-form Entity.
// This is a lower-level API and you probably actually want the higher-level entities package.
type Entity struct {
	// Snowflake ID, separate from the public one, used for foreign key consistency.
	ID snowflake.ID `json:"_id"`

	// Raw JSON data.
	Data JSONB `json:"_data"`

	// "Kind" of an entity
	// Kinds determine what special server side behavior applies
	Kind string `json:"_kind"`

	// Concrete Object representation of this Entity.
	// Modifications to this Object will be written back (and synched to Data) by EntityStore.Save().
	obj *ld.Object `gorm:"-"`
}

func NewEntity(kind string, data []byte) (*Entity, error) {
	id, err := lib.GenSnowflake(config.NodeID())
	return &Entity{ID: id, Data: data, Kind: kind}, err
}

// Lazily constructs and caches an Object from e.Data.
func (e *Entity) Object() (*ld.Object, error) {
	if e.obj == nil {
		obj, err := ld.NewObject(e.Data)
		if err != nil {
			return nil, err
		}
		e.obj = obj
	}
	return e.obj, nil
}

// EntityStore stores Entities in their raw database form.
type EntityStore interface {
	// GetBySnowflake returns an Entity by its snowflake, eg. "353894652568535040".
	GetBySnowflake(id snowflake.ID) (*Entity, error)

	// GetByID returns an Entity by its ID, eg. "https://example.com/@johnsmith.
	GetByID(id string) (*Entity, error)

	// Save stores an Entity using an upsert. Updates Data if the object has been modified.
	Save(e *Entity) error
}

type entityStore struct {
	DB *gorm.DB
}

func NewEntityStore(db *gorm.DB) EntityStore {
	return &entityStore{db}
}

func (s entityStore) GetBySnowflake(id snowflake.ID) (*Entity, error) {
	var e Entity
	return &e, s.DB.First(&e, Entity{ID: id}).Error
}

func (s entityStore) GetByID(id string) (*Entity, error) {
	var e Entity
	return &e, s.DB.First(&e, `data->>'@id' = ?`, id).Error
}

func (s entityStore) Save(e *Entity) error {
	if obj := e.obj; obj != nil {
		data, err := json.Marshal(obj)
		if err != nil {
			return errors.Wrap(err, "couldn't serialise back")
		}
		e.Data = data
	}
	return s.DB.Set(gormInsertOption, entityOnConflict).Create(e).Error
}
