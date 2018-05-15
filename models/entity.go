package models

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jinzhu/gorm"

	"github.com/liclac/meow/config"
	"github.com/liclac/meow/lib"
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
}

func NewEntity(kind string, data []byte) (*Entity, error) {
	id, err := lib.GenSnowflake(config.NodeID())
	return &Entity{ID: id, Data: data, Kind: kind}, err
}

// EntityStore stores Entities in their raw database form.
// This is a lower-level API and you probably actually want the higher-level entities package.
type EntityStore interface {
	// GetBySnowflake returns an Entity by its snowflake, eg. "353894652568535040".
	GetBySnowflake(id snowflake.ID) (*Entity, error)

	// GetByID returns an Entity by its ID, eg. "https://example.com/@johnsmith.
	GetByID(id string) (*Entity, error)

	// Save stores an Entity using an upsert.
	Save(e Entity) error
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

func (s entityStore) Save(e Entity) error {
	return s.DB.Set(gormInsertOption, entityOnConflict).Create(&e).Error
}
