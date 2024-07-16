// Code generated by MockGen. DO NOT EDIT.
// Source: ../db.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIDB is a mock of IDB interface.
type MockIDB struct {
	ctrl     *gomock.Controller
	recorder *MockIDBMockRecorder
}

// MockIDBMockRecorder is the mock recorder for MockIDB.
type MockIDBMockRecorder struct {
	mock *MockIDB
}

// NewMockIDB creates a new mock instance.
func NewMockIDB(ctrl *gomock.Controller) *MockIDB {
	mock := &MockIDB{ctrl: ctrl}
	mock.recorder = &MockIDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDB) EXPECT() *MockIDBMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockIDB) Ping(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockIDBMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockIDB)(nil).Ping), arg0)
}