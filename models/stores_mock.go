package models

import (
	"github.com/golang/mock/gomock"
)

// Unlike the other stores, we're not autogenerating the mock for Stores, because it's more useful
// to have a Store that creates Mock<*>Store instances that we can run tests on than to have to go
// mockStore.EXPECT().Clients().Returns(myExistingMockClientStore) etc. before every real EXPECT().
// That would also break in obnoxious ways if tested functions change the way they reuse stores,
// which should not be part of the API contract in the slightest.
type MockStores struct {
	ctrl *gomock.Controller

	EntityStore        *MockEntityStore
	UserStore          *MockUserStore
	ClientStore        *MockClientStore
	AuthorizationStore *MockAuthorizationStore
	AccessTokenStore   *MockAccessTokenStore
	RefreshTokenStore  *MockRefreshTokenStore
	StreamItemStore    *MockStreamItemStore
}

func NewMockStores(ctrl *gomock.Controller) *MockStores {
	return &MockStores{ctrl: ctrl}
}

func (s *MockStores) Entities() EntityStore {
	if s.EntityStore == nil {
		s.EntityStore = NewMockEntityStore(s.ctrl)
	}
	return s.EntityStore
}

func (s *MockStores) Users() UserStore {
	if s.UserStore == nil {
		s.UserStore = NewMockUserStore(s.ctrl)
	}
	return s.UserStore
}

func (s *MockStores) Clients() ClientStore {
	if s.ClientStore == nil {
		s.ClientStore = NewMockClientStore(s.ctrl)
	}
	return s.ClientStore
}

func (s *MockStores) Authorizations() AuthorizationStore {
	if s.AuthorizationStore == nil {
		s.AuthorizationStore = NewMockAuthorizationStore(s.ctrl)
	}
	return s.AuthorizationStore
}

func (s *MockStores) AccessTokens() AccessTokenStore {
	if s.AccessTokenStore == nil {
		s.AccessTokenStore = NewMockAccessTokenStore(s.ctrl)
	}
	return s.AccessTokenStore
}

func (s *MockStores) RefreshTokens() RefreshTokenStore {
	if s.RefreshTokenStore == nil {
		s.RefreshTokenStore = NewMockRefreshTokenStore(s.ctrl)
	}
	return s.RefreshTokenStore
}

func (s *MockStores) StreamItems() StreamItemStore {
	if s.StreamItemStore == nil {
		s.StreamItemStore = NewMockStreamItemStore(s.ctrl)
	}
	return s.StreamItemStore
}
