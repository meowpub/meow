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
}

// NewStore creates a new Entity Store on a raw EntityStore
func NewStore(rawStore models.EntityStore) *Store {
	return &Store{
		rawStore:        rawStore,
		liveByID:        make(map[string]Entity),
		liveBySnowflake: make(map[snowflake.ID]Entity),
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

	e, err := kind.Unmarshall(raw.ID, map_)
	if err != nil {
		return nil, errors.Wrap(err, "Unmarshalling entity")
	}

	if e.GetSnowflake() != raw.ID {
		return nil, errors.New("Unmarshalling entity produced wrong snowflake")
	}

	s.liveByID[e.GetID()] = e
	s.liveBySnowflake[raw.ID] = e

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

// NewEntity creates a new Entity
func (s *Store) NewEntity(ctx context.Context, kind string, id string) (Entity, error) {
	k, err := GetKind(kind)
	if err != nil {
		return nil, errors.Wrap(err, "Getting kind")
	}

	flake, err := lib.GenSnowflake(config.NodeID())
	if err != nil {
		return nil, errors.Wrap(err, "Error generating snowflake while creating entity")
	}

	n, err := k.New(flake, id)
	if err != nil {
		return nil, errors.Wrap(err, "Creating new entity")
	}

	if err := s.insert(n); err != nil {
		return nil, err
	}

	return n, n.Init(ctx)
}

// NewChildEntity creates a child Entity
// (will rename on conflict)
func (s *Store) NewChildEntity(ctx context.Context, parent Entity, kind string, suggestedSuffix string) (Entity, error) {
	k, err := GetKind(kind)
	if err != nil {
		return nil, errors.Wrap(err, "Getting kind")
	}

	flake, err := lib.GenSnowflake(config.NodeID())
	if err != nil {
		return nil, errors.Wrap(err, "Error generating snowflake while creating entity")
	}

	id := parent.GetID() + "/" + suggestedSuffix

	if suggestedSuffix != "" {
		_, err := s.GetByID(id)
		if err != nil {
			if !models.IsNotFound(err) {
				return nil, errors.Wrap(err, "While discovering if entity ID conflicts")
			}
			// Fallthrough
		} else {
			// We found the thing, so uniquify with our snowflake
			id = id + flake.String()
		}
	} else {
		// No suggested suffix so just use our snowflake
		id = id + flake.String()
	}

	n, err := k.New(flake, id)
	if err != nil {
		return nil, errors.Wrap(err, "Creating new entity")
	}

	if err := s.insert(n); err != nil {
		return nil, err
	}

	return n, n.Init(ctx)
}

// Insert inserts a new entity
func (s *Store) insert(e Entity) error {
	if _, ok := s.liveByID[e.GetID()]; ok {
		return errors.Errorf("Attempt to insert duplicate entity '%s' into store", e.GetID())
	}

	flake := e.GetSnowflake()
	if flake == 0 {
		return errors.New("Attempt to insert entity with 0 snowflake!")
	} else if s.liveBySnowflake[flake] != nil {
		return errors.New("Attempt to insert entity with duplicate snowflake")
	}

	s.liveBySnowflake[flake] = e
	s.liveByID[e.GetID()] = e

	return nil
}

// Save saves an entity into the database
func (s *Store) Save(e Entity) error {
	// Dig out this entity by its snowflake because we need to upcast
	// else we may accidentally slice the object in half when marshalling
	e = s.liveBySnowflake[e.GetSnowflake()]

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
		ID:   e.GetSnowflake(),
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
