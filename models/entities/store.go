package entities

import (
	"context"
	"encoding/json"

	"github.com/bwmarrin/snowflake"
	"github.com/pkg/errors"

	"github.com/meowpub/meow/config"
	"github.com/meowpub/meow/lib"
	"github.com/meowpub/meow/models"
)

type key int

var storeKey key

type Store struct {
	rawStore        models.EntityStore
	liveByID        map[string]Entity
	liveBySnowflake map[snowflake.ID]Entity
	liveToSnowflake map[Entity]snowflake.ID
}

// NewStore creates a new Entity Store on a raw EntityStore
func NewStore(rawStore models.EntityStore) *Store {
	return &Store{
		rawStore:        rawStore,
		liveByID:        make(map[string]Entity),
		liveBySnowflake: make(map[snowflake.ID]Entity),
		liveToSnowflake: make(map[Entity]snowflake.ID),
	}
}

// inflateEntity converts a raw Entity into an Entity
func (s *Store) inflateEntity(raw *models.Entity) (Entity, error) {
	kind, err := GetKind(raw.Kind)
	if err != nil {
		return nil, err
	}

	var map_ map[string]interface{}
	if err := json.Unmarshal(raw.Data, &map_); err != nil {
		return nil, errors.Wrap(err, "Unmarshalling JSON data")
	}

	e, err := kind.Unmarshall(map_)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshalling entity")
	}

	s.liveByID[e.GetID()] = e
	s.liveBySnowflake[raw.ID] = e
	s.liveToSnowflake[e] = raw.ID
	e.SetSnowflake(raw.ID)

	return e, nil
}

// GetBySnowflake gets an Entity by a Snowflake
func (s *Store) GetBySnowflake(id snowflake.ID) (Entity, error) {
	if entity, ok := s.liveBySnowflake[id]; ok {
		return entity, nil
	}

	if rawEntity, err := s.rawStore.GetBySnowflake(id); err == nil {
		return s.inflateEntity(rawEntity)
	} else {
		return nil, errors.Wrapf(err, "Unable to find entity by snowflake %s", id)
	}
}

// GetByID gets an entity by its (URI) ID
func (s *Store) GetByID(id string) (Entity, error) {
	if entity, ok := s.liveByID[id]; ok {
		return entity, nil
	}

	if rawEntity, err := s.rawStore.GetByID(id); err == nil {
		return s.inflateEntity(rawEntity)
	} else {
		return nil, errors.Wrapf(err, "Unable to find '%s'", id)
	}
}

// Insert inserts a new entity
func (s *Store) Insert(e Entity) error {
	if _, ok := s.liveToSnowflake[e]; ok {
		return errors.New("Attempt to insert entity into Store which is already inserted")
	}

	if _, ok := s.liveByID[e.GetID()]; ok {
		return errors.Errorf("Attempt to insert duplicate entity '%s' into store", e.GetID())
	}

	flake := e.GetSnowflake()
	if flake == 0 {
		var err error
		flake, err = lib.GenSnowflake(config.NodeID())
		if err != nil {
			return errors.Wrap(err, "While inserting entity")
		}
	} else {
		if s.liveBySnowflake[flake] != nil {
			return errors.New("Attempt to insert entity with duplicate snowflake")
		}
	}

	s.liveToSnowflake[e] = flake
	s.liveBySnowflake[flake] = e
	s.liveByID[e.GetID()] = e
	e.SetSnowflake(flake)

	return nil
}

// Save saves an entity into the database
func (s *Store) Save(e Entity) error {
	flake, ok := s.liveToSnowflake[e]
	if !ok {
		return errors.New("Attempt to save live entity which doesn't come from this store")
	}

	kind := e.GetKind()

	map_, err := kind.Marshall(e)
	if err != nil {
		return errors.Wrap(err, "Marhalling entity for save")
	}

	data, err := json.Marshal(map_)
	if err != nil {
		return errors.Wrap(err, "Marshalling entity to JSON for save")
	}

	r := models.Entity{
		ID:   flake,
		Kind: kind.Name,
		Data: data,
	}

	err = s.rawStore.Save(r)
	if err != nil {
		return errors.Wrap(err, "Saving entity")
	} else {
		return nil
	}
}

// WithStore creates a new context with this entity Store
func WithStore(ctx context.Context, st *Store) context.Context {
	return context.WithValue(ctx, storeKey, st)
}

// GetStore gets the entity Store from a Context
func GetStore(ctx context.Context) *Store {
	store, _ := ctx.Value(storeKey).(*Store)
	return store
}
