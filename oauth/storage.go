package oauth

import (
	"time"

	"github.com/RangelReale/osin"
	"github.com/liclac/meow/models"
)

// Storage implements osin's Storage interface. Clients go in a database (see models.go),
// tokens in Redis.
type Storage struct {
	Clients models.ClientStore
	Auths   models.AuthorizationStore
}

// NewStorage returns a new Storage instance.
func NewStorage(clients models.ClientStore, auths models.AuthorizationStore) *Storage {
	return &Storage{
		Clients: clients,
		Auths:   auths,
	}
}

// Clone just returns itself.
func (s *Storage) Clone() osin.Storage {
	return s
}

// Close closes any resources attached to the store, in this case nothing.
func (s Storage) Close() {
}

// GetClient returns the client with the given ID.
func (s Storage) GetClient(id string) (osin.Client, error) {
	cl, err := s.Clients.Get(id)
	if models.IsNotFound(err) {
		return nil, osin.ErrNotFound
	}
	return cl, err
}

// SaveAuthorize stores authorization data in Redis for later use, w/ a TTL.
func (s Storage) SaveAuthorize(auth *osin.AuthorizeData) error {
	return s.Auths.Set(&models.Authorization{
		Code:        auth.Code,
		ClientID:    auth.Client.GetId(),
		Scope:       auth.Scope,
		RedirectURI: auth.RedirectUri,
		State:       auth.State,
	}, time.Duration(auth.ExpiresIn)*time.Second)
}

// LoadAuthorize loads previously saved authorization data from Redis.
// The Client field must be populated.
func (s Storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	auth, err := s.Auths.Get(code)
	if err != nil {
		if models.IsNotFound(err) {
			return nil, osin.ErrNotFound
		}
		return nil, err
	}
	client, err := s.GetClient(auth.ClientID)
	if err != nil {
		return nil, err
	}
	return &osin.AuthorizeData{
		Client:      client,
		Code:        code,
		Scope:       auth.Scope,
		RedirectUri: auth.RedirectURI,
		State:       auth.State,
	}, nil
}

// RemoveAuthorize explicitly removes authorization data.
func (s Storage) RemoveAuthorize(code string) error {
	return s.Auths.Delete(code)
}

// SaveAccess saves the refresh token to redis.
// It's possible for there not to be an access token; NOP in that case.
func (s Storage) SaveAccess(access *osin.AccessData) error {
	panic("not implemented")
}

// LoadAccess looks up an access token in Redis. The Client field must be populated, but
// AuthorizeData and AccessData do NOT need to be loaded if it's not easily available.
func (s Storage) LoadAccess(token string) (*osin.AccessData, error) {
	panic("not implemented")
}

// RemoveAccess removes the token from Redis.
func (s Storage) RemoveAccess(token string) error {
	panic("not implemented")
}

// LoadRefresh loads access data out of redis. The Client field must be populated, but
// AuthorizeData and AccessData do NOT need to be loaded if it's not easily available.
func (s Storage) LoadRefresh(token string) (*osin.AccessData, error) {
	panic("not implemented")
}

// RemoveRefresh deletes a refresh token from redis.
func (s Storage) RemoveRefresh(token string) error {
	panic("not implemented")
}
