package models

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

// Stores is an interface that returns other Stores. There are too many of them at this point to
// keep them all around individually. I wasn't allowed to name this the StoreStore :(
type Stores interface {
	Entities() EntityStore
	Users() UserStore
	Clients() ClientStore
	Authorizations() AuthorizationStore
	AccessTokens() AccessTokenStore
	RefreshTokens() RefreshTokenStore
}

type stores struct {
	DB *gorm.DB
	R  *redis.Client
	K  string

	entityStore        EntityStore
	userStore          UserStore
	clientStore        ClientStore
	authorizationStore AuthorizationStore
	accessTokenStore   AccessTokenStore
	refreshTokenStore  RefreshTokenStore
}

func NewStores(db *gorm.DB, r *redis.Client, keyspace string) Stores {
	return &stores{
		DB: db,
		R:  r,
		K:  keyspace,
	}
}

func (s *stores) Entities() EntityStore {
	if s.entityStore == nil {
		s.entityStore = NewEntityStore(s.DB)
	}
	return s.entityStore
}

func (s *stores) Users() UserStore {
	if s.userStore == nil {
		s.userStore = NewUserStore(s.DB)
	}
	return s.userStore
}

func (s *stores) Clients() ClientStore {
	if s.clientStore == nil {
		s.clientStore = NewClientStore(s.DB)
	}
	return s.clientStore
}

func (s *stores) Authorizations() AuthorizationStore {
	if s.authorizationStore == nil {
		s.authorizationStore = NewAuthorizationStore(s.K, s.R)
	}
	return s.authorizationStore
}

func (s *stores) AccessTokens() AccessTokenStore {
	if s.accessTokenStore == nil {
		s.accessTokenStore = NewAccessTokenStore(s.K, s.R)
	}
	return s.accessTokenStore
}

func (s *stores) RefreshTokens() RefreshTokenStore {
	if s.refreshTokenStore == nil {
		s.refreshTokenStore = NewRefreshTokenStore(s.K, s.R)
	}
	return s.refreshTokenStore
}
