package models

import (
	"encoding/json"

	"github.com/go-redis/redis"
)

//go:generate mockgen -package=models -source=refreshtoken.go -destination=refreshtoken.mock.go

// RefreshToken stores a persistent refresh token in Redis. See also: AccessToken.
type RefreshToken struct {
	Token    string `json:"token"`
	ClientID string `json:"client_id"`
	Scope    string `json:"scope"`
	AuthorizationUserData
}

type RefreshTokenStore interface {
	Set(token *RefreshToken) error
	Get(token string) (*RefreshToken, error)
	Delete(token string) error
}

type refreshTokenStore struct {
	K string
	R *redis.Client
}

func NewRefreshTokenStore(keyspace string, r *redis.Client) RefreshTokenStore {
	return &refreshTokenStore{keyspace, r}
}

func (s *refreshTokenStore) key(code string) string {
	return s.K + ":oauth2:refresh_tokens:" + code
}

func (s *refreshTokenStore) Set(token *RefreshToken) error {
	data, err := json.Marshal(token)
	if err != nil {
		return err
	}
	return s.R.Set(s.key(token.Token), data, 0).Err()
}

func (s *refreshTokenStore) Get(token string) (*RefreshToken, error) {
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
	var tok RefreshToken
	return &tok, json.Unmarshal(data, &tok)
}

func (s *refreshTokenStore) Delete(token string) error {
	return s.R.Del(s.key(token)).Err()
}
