package oauth

import (
	"strconv"
	"time"

	"github.com/RangelReale/osin"
	"github.com/bwmarrin/snowflake"
	"github.com/liclac/meow/models"
	"go.uber.org/multierr"
)

// Storage implements osin's Storage interface. Clients go in a database (see models.go),
// tokens in Redis.
type Storage struct {
	Stores models.Stores
}

// NewStorage returns a new Storage instance.
func NewStorage(stores models.Stores) *Storage {
	return &Storage{stores}
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
	flake, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	cl, err := s.Stores.Clients().Get(snowflake.ID(flake))
	if models.IsNotFound(err) {
		return nil, osin.ErrNotFound
	}
	return cl, err
}

// SaveAuthorize stores authorization data in Redis for later use, w/ a TTL.
func (s Storage) SaveAuthorize(auth *osin.AuthorizeData) error {
	return s.Stores.Authorizations().Set(&models.Authorization{
		Code:        auth.Code,
		ClientID:    auth.Client.GetId(),
		Scope:       auth.Scope,
		RedirectURI: auth.RedirectUri,
		State:       auth.State,
		AuthorizationUserData: auth.UserData.(models.AuthorizationUserData),
	}, time.Duration(auth.ExpiresIn)*time.Second)
}

// LoadAuthorize loads previously saved authorization data from Redis.
// The Client field must be populated.
func (s Storage) LoadAuthorize(code string) (*osin.AuthorizeData, error) {
	auth, err := s.Stores.Authorizations().Get(code)
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
		UserData:    auth.AuthorizationUserData,
	}, nil
}

// RemoveAuthorize explicitly removes authorization data.
func (s Storage) RemoveAuthorize(code string) error {
	return s.Stores.Authorizations().Delete(code)
}

// SaveAccess saves the access- and refresh tokens to redis.
// This may be called without one or the other, and this should be handled gracefully.
func (s Storage) SaveAccess(access *osin.AccessData) error {
	userData := access.UserData.(models.AuthorizationUserData)

	var errs []error
	if access.AccessToken != "" {
		errs = append(errs, s.Stores.AccessTokens().Set(&models.AccessToken{
			Token:    access.AccessToken,
			ClientID: access.Client.GetId(),
			Scope:    access.Scope,
			AuthorizationUserData: userData,
		}, time.Duration(access.ExpiresIn)*time.Second))
	}
	if access.RefreshToken != "" {
		errs = append(errs, s.Stores.RefreshTokens().Set(&models.RefreshToken{
			Token:    access.AccessToken,
			ClientID: access.Client.GetId(),
			Scope:    access.Scope,
			AuthorizationUserData: userData,
		}))
	}
	return multierr.Combine(errs...)
}

// LoadAccess looks up an access token in Redis. The Client field must be populated, but
// AuthorizeData and AccessData do NOT need to be loaded if it's not easily available.
func (s Storage) LoadAccess(token string) (*osin.AccessData, error) {
	tok, err := s.Stores.AccessTokens().Get(token)
	if err != nil {
		return nil, err
	}
	client, err := s.GetClient(tok.ClientID)
	if err != nil {
		return nil, err
	}
	return &osin.AccessData{
		Client:      client,
		AccessToken: tok.Token,
		Scope:       tok.Scope,
		UserData:    tok.AuthorizationUserData,
	}, nil
}

// RemoveAccess removes the token from Redis.
func (s Storage) RemoveAccess(token string) error {
	return s.Stores.AccessTokens().Delete(token)
}

// LoadRefresh loads access data out of redis. The Client field must be populated, but
// AuthorizeData and AccessData do NOT need to be loaded if it's not easily available.
func (s Storage) LoadRefresh(token string) (*osin.AccessData, error) {
	tok, err := s.Stores.RefreshTokens().Get(token)
	if err != nil {
		return nil, err
	}
	client, err := s.GetClient(tok.ClientID)
	if err != nil {
		return nil, err
	}
	return &osin.AccessData{
		Client:       client,
		RefreshToken: tok.Token,
		Scope:        tok.Scope,
		UserData:     tok.AuthorizationUserData,
	}, nil
}

// RemoveRefresh deletes a refresh token from redis.
func (s Storage) RemoveRefresh(token string) error {
	return s.Stores.RefreshTokens().Delete(token)
}
