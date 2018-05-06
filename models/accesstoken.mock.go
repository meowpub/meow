// Code generated by MockGen. DO NOT EDIT.
// Source: accesstoken.go

// Package models is a generated GoMock package.
package models

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockAccessTokenStore is a mock of AccessTokenStore interface
type MockAccessTokenStore struct {
	ctrl     *gomock.Controller
	recorder *MockAccessTokenStoreMockRecorder
}

// MockAccessTokenStoreMockRecorder is the mock recorder for MockAccessTokenStore
type MockAccessTokenStoreMockRecorder struct {
	mock *MockAccessTokenStore
}

// NewMockAccessTokenStore creates a new mock instance
func NewMockAccessTokenStore(ctrl *gomock.Controller) *MockAccessTokenStore {
	mock := &MockAccessTokenStore{ctrl: ctrl}
	mock.recorder = &MockAccessTokenStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccessTokenStore) EXPECT() *MockAccessTokenStoreMockRecorder {
	return m.recorder
}

// Set mocks base method
func (m *MockAccessTokenStore) Set(token *AccessToken, ttl time.Duration) error {
	ret := m.ctrl.Call(m, "Set", token, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set
func (mr *MockAccessTokenStoreMockRecorder) Set(token, ttl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockAccessTokenStore)(nil).Set), token, ttl)
}

// Get mocks base method
func (m *MockAccessTokenStore) Get(token string) (*AccessToken, error) {
	ret := m.ctrl.Call(m, "Get", token)
	ret0, _ := ret[0].(*AccessToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAccessTokenStoreMockRecorder) Get(token interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccessTokenStore)(nil).Get), token)
}

// Delete mocks base method
func (m *MockAccessTokenStore) Delete(token string) error {
	ret := m.ctrl.Call(m, "Delete", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockAccessTokenStoreMockRecorder) Delete(token interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccessTokenStore)(nil).Delete), token)
}