// Code generated by MockGen. DO NOT EDIT.
// Source: snowflakes.go

// Package lib is a generated GoMock package.
package lib

import (
	snowflake "github.com/bwmarrin/snowflake"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockSnowflakeGenerator is a mock of SnowflakeGenerator interface
type MockSnowflakeGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockSnowflakeGeneratorMockRecorder
}

// MockSnowflakeGeneratorMockRecorder is the mock recorder for MockSnowflakeGenerator
type MockSnowflakeGeneratorMockRecorder struct {
	mock *MockSnowflakeGenerator
}

// NewMockSnowflakeGenerator creates a new mock instance
func NewMockSnowflakeGenerator(ctrl *gomock.Controller) *MockSnowflakeGenerator {
	mock := &MockSnowflakeGenerator{ctrl: ctrl}
	mock.recorder = &MockSnowflakeGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSnowflakeGenerator) EXPECT() *MockSnowflakeGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method
func (m *MockSnowflakeGenerator) Generate() snowflake.ID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(snowflake.ID)
	return ret0
}

// Generate indicates an expected call of Generate
func (mr *MockSnowflakeGeneratorMockRecorder) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockSnowflakeGenerator)(nil).Generate))
}