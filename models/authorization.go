package models

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

//go:generate mockgen -package=models -source=authorization.go -destination=authorization.mock.go

// Authorization stores server-side state for an authorization. It's stored temporarily, before
// being traded for an AccessToken and a RefreshToken and subsequently deleted.
type Authorization struct {
	ClientID    string `json:"client_id"`
	Scope       string `json:"scope"`
	RedirectURI string `json:"redirect_uri"`
	State       string `json:"state"`
}

// AuthorizationStore stores Authorizations (in Redis).
type AuthorizationStore interface {
	Set(code string, auth *Authorization, ttl time.Duration) error
	Get(code string) (*Authorization, error)
	Delete(code string) error
}

type authorizationStore struct {
	K string
	R *redis.Client
}

func NewAuthorizationStore(keyspace string, r *redis.Client) AuthorizationStore {
	return &authorizationStore{keyspace, r}
}

func (s authorizationStore) key(code string) string {
	return s.K + ":oauth2:authorizations:" + code
}

func (s authorizationStore) Set(code string, auth *Authorization, ttl time.Duration) error {
	if ttl == 0 {
		return ErrNoTTL
	}
	data, err := json.Marshal(auth)
	if err != nil {
		return err
	}
	return s.R.Set(s.key(code), data, ttl).Err()
}

func (s authorizationStore) Get(code string) (*Authorization, error) {
	cmd := s.R.Get(s.key(code))
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return nil, notFoundError("invalid authorization code")
		}
		return nil, err
	}
	data, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	var auth Authorization
	return &auth, json.Unmarshal(data, &auth)
}

func (s authorizationStore) Delete(code string) error {
	return s.R.Del(s.key(code)).Err()
}
