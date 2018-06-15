package models

import (
	"github.com/golang/mock/gomock"
)

// Unlike the other stores, we're not autogenerating the mock for Stores, because it's more useful
// to have a Store that creates Mock<*>Store instances that we can run tests on than to have to go
// mockStore.EXPECT().Clients().Returns(myExistingMockClientStore) etc. before every real EXPECT().
// That would also break in obnoxious ways if tested functions change the way they reuse stores,
// which should not be part of the API contract in the slightest.
type mockStores struct {
	ctrl *gomock.Controller

	entityStore        EntityStore
	userStore          UserStore
	clientStore        ClientStore
	authorizationStore AuthorizationStore
	accessTokenStore   AccessTokenStore
	refreshTokenStore  RefreshTokenStore
	streamItemStore    StreamItemStore
}

func NewMockStores(ctrl *gomock.Controller) Stores {
	return &mockStores{ctrl: ctrl}
}

func (s *mockStores) Entities() EntityStore {
	if s.entityStore == nil {
		s.entityStore = NewMockEntityStore(s.ctrl)
	}
	return s.entityStore
}

func (s *mockStores) Users() UserStore {
	if s.userStore == nil {
		s.userStore = NewMockUserStore(s.ctrl)
	}
	return s.userStore
}

func (s *mockStores) Clients() ClientStore {
	if s.clientStore == nil {
		s.clientStore = NewMockClientStore(s.ctrl)
	}
	return s.clientStore
}

func (s *mockStores) Authorizations() AuthorizationStore {
	if s.authorizationStore == nil {
		s.authorizationStore = NewMockAuthorizationStore(s.ctrl)
	}
	return s.authorizationStore
}

func (s *mockStores) AccessTokens() AccessTokenStore {
	if s.accessTokenStore == nil {
		s.accessTokenStore = NewMockAccessTokenStore(s.ctrl)
	}
	return s.accessTokenStore
}

func (s *mockStores) RefreshTokens() RefreshTokenStore {
	if s.refreshTokenStore == nil {
		s.refreshTokenStore = NewMockRefreshTokenStore(s.ctrl)
	}
	return s.refreshTokenStore
}

func (s *mockStores) StreamItems() StreamItemStore {
	if s.streamItemStore == nil {
		s.streamItemStore = NewMockStreamItemStore(s.ctrl)
	}
	return s.streamItemStore
}
