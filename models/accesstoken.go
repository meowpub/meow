package models

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

// AccessToken stores an active access token in Redis. See also: RefreshToken.
type AccessToken struct {
	Token    string `json:"token"`
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
	AuthorizationUserData
}

type AccessTokenStore interface {
	Set(token *AccessToken, ttl time.Duration) error
	Get(token string) (*AccessToken, error)
	Delete(token string) error
}

type accessTokenStore struct {
	K string
	R *redis.Client
}

func NewAccessTokenStore(keyspace string, r *redis.Client) AccessTokenStore {
	return &accessTokenStore{keyspace, r}
}

func (s *accessTokenStore) key(code string) string {
	return s.K + ":oauth2:access_tokens:" + code
}

func (s *accessTokenStore) Set(token *AccessToken, ttl time.Duration) error {
	if ttl == 0 {
		return ErrNoTTL
	}
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return s.R.Set(s.key(token.Token), data, ttl).Err()
}

func (s *accessTokenStore) Get(token string) (*AccessToken, error) {
	cmd := s.R.Get(s.key(token))
	if err := cmd.Err(); err != nil {
		if err == redis.Nil {
			return nil, NotFoundError("token is invalid or expired")
		}
		return nil, err
	}
	data, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	var tok AccessToken
	return &tok, json.Unmarshal(data, &tok)
}

func (s *accessTokenStore) Delete(token string) error {
	return s.R.Del(s.key(token)).Err()
}
