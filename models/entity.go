package models

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"
	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
)

//go:generate mockgen -package=models -source=entity.go -destination=entity.mock.go

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
}

func NewEntity(data []byte) (*Entity, error) {
	id, err := lib.GenSnowflake(config.NodeID())
	return &Entity{ID: id, Data: data}, err
}

// EntityStore stores Entities in their raw database form.
// This is a lower-level API and you probably actually want the higher-level entities package.
type EntityStore interface {
	// GetBySnowflake returns an Entity by its snowflake, eg. "353894652568535040".
	GetBySnowflake(id snowflake.ID) (*Entity, error)

	// GetByID returns an Entity by its ID, eg. "https://example.com/@johnsmith.
	GetByID(id string) (*Entity, error)

	// Save stores an Entity using an upsert.
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
	// Note about this expression: we are NOT conflicting on the snowflake.
	// These are the possible scenarios at work here:
	//
	// 0. Insert with no snowflake -> incorrect call, error out.
	// 1. Insert with a unique snowflake and ID -> no conflict.
	// 2. Insert with a reused snowflake and unique ID -> this is either an incorrect Save call or
	//    we just got hit by the birthday paradox, error the heck out.
	// 3. Insert with a unique snowflake and reused ID -> discard the new snowflake and reuse the
	//    old one; this preserves foreign key integrity without forcing the models/entities API to
	//    hang onto snowflakes.
	return s.DB.Set(gormInsertOption, `ON CONFLICT ((data->>'@id')) DO UPDATE SET data = EXCLUDED.data`).Create(e).Error
}
