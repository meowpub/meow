// Code generated by MockGen. DO NOT EDIT.
// Source: entity.go

// Package models is a generated GoMock package.
package models

import (
	snowflake "github.com/bwmarrin/snowflake"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEntityStore is a mock of EntityStore interface
type MockEntityStore struct {
	ctrl     *gomock.Controller
	recorder *MockEntityStoreMockRecorder
}

// MockEntityStoreMockRecorder is the mock recorder for MockEntityStore
type MockEntityStoreMockRecorder struct {
	mock *MockEntityStore
}

// NewMockEntityStore creates a new mock instance
func NewMockEntityStore(ctrl *gomock.Controller) *MockEntityStore {
	mock := &MockEntityStore{ctrl: ctrl}
	mock.recorder = &MockEntityStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEntityStore) EXPECT() *MockEntityStoreMockRecorder {
	return m.recorder
}

// GetBySnowflake mocks base method
func (m *MockEntityStore) GetBySnowflake(id snowflake.ID) (*Entity, error) {
	ret := m.ctrl.Call(m, "GetBySnowflake", id)
	ret0, _ := ret[0].(*Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySnowflake indicates an expected call of GetBySnowflake
func (mr *MockEntityStoreMockRecorder) GetBySnowflake(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySnowflake", reflect.TypeOf((*MockEntityStore)(nil).GetBySnowflake), id)
}

// GetByID mocks base method
func (m *MockEntityStore) GetByID(id string) (*Entity, error) {
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(*Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockEntityStoreMockRecorder) GetByID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockEntityStore)(nil).GetByID), id)
}

// Save mocks base method
func (m *MockEntityStore) Save(e *Entity) error {
	ret := m.ctrl.Call(m, "Save", e)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockEntityStoreMockRecorder) Save(e interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockEntityStore)(nil).Save), e)
}
