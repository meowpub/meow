package models

import (
	"encoding/json"

	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/ld"
	"github.com/meowpub/meow/lib"
)

//go:generate mockgen -package=models -source=entity.go -destination=entity.mock.go

var entityOnConflict = genOnConflict(Entity{}, "(data->>'@id')", "id")

// Kind of an Entity.
type EntityKind string

const (
	ObjectEntity = "object" // Default Kind of an Entity.
)

// An Entity represents a raw, database-form Entity.
// This is a lower-level API and you probably actually want the higher-level entities package.
type Entity struct {
	// Snowflake ID, separate from the public one, used for foreign key consistency.
	ID snowflake.ID `json:"_id"`

	// Raw JSON data. You probably want to use Obj instead.
	Data JSONB `json:"_data"`

	// "Kind" of an entity, normally ObjectEntity.
	// Kinds determine what special server side behavior applies.
	Kind EntityKind `json:"_kind"`

	// Concrete Object representation of this Entity.
	// Modifications to this Object will be written back (and synched to Data) by EntityStore.Save().
	Obj *ld.Object `json:"_obj" gorm:"-"`
}

func NewEntity(kind EntityKind, data []byte) (*Entity, error) {
	id, err := lib.GenSnowflake(config.NodeID())
	if err != nil {
		return nil, err
	}
	e := &Entity{ID: id, Data: data, Kind: kind}
	return e, e.SyncDataToObject()
}

// Overwrites Object with Data. Called automatically by EntityStore.Get*() and NewEntity().
func (e *Entity) SyncDataToObject() error {
	if len(e.Data) > 0 {
		obj, err := ld.ParseObject(e.Data)
		if err != nil {
			return err
		}
		e.Obj = obj
	} else {
		e.Obj = nil
	}
	return nil
}

// Overwrites Data with Object. Called automatically by EntityStore.Save().
func (e *Entity) SyncObjectToData() error {
	if e.Obj != nil {
		data, err := json.Marshal(e.Obj)
		if err != nil {
			return err
		}
		e.Data = data
	}
	return nil
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
	if err := s.DB.First(&e, Entity{ID: id}).Error; err != nil {
		return &e, err
	}
	return &e, e.SyncDataToObject()
}

func (s entityStore) GetByID(id string) (*Entity, error) {
	var e Entity
	if err := s.DB.First(&e, `data->>'@id' = ?`, id).Error; err != nil {
		return &e, err
	}
	return &e, e.SyncDataToObject()
}

func (s entityStore) Save(e *Entity) error {
	if err := e.SyncObjectToData(); err != nil {
		return err
	}
	return s.DB.Set(gormInsertOption, entityOnConflict).Create(e).Error
}
