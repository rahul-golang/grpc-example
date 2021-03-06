// Code generated by MockGen. DO NOT EDIT.
// Source: downstream/network_downstream.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "interview/models"
	reflect "reflect"
)

// MockNetworkDownStream is a mock of NetworkDownStream interface
type MockNetworkDownStream struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkDownStreamMockRecorder
}

// MockNetworkDownStreamMockRecorder is the mock recorder for MockNetworkDownStream
type MockNetworkDownStreamMockRecorder struct {
	mock *MockNetworkDownStream
}

// NewMockNetworkDownStream creates a new mock instance
func NewMockNetworkDownStream(ctrl *gomock.Controller) *MockNetworkDownStream {
	mock := &MockNetworkDownStream{ctrl: ctrl}
	mock.recorder = &MockNetworkDownStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkDownStream) EXPECT() *MockNetworkDownStreamMockRecorder {
	return m.recorder
}

// ExecuteNetworkRequest mocks base method
func (m *MockNetworkDownStream) ExecuteNetworkRequest(ctx context.Context, networkKey int64) (models.UserKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteNetworkRequest", ctx, networkKey)
	ret0, _ := ret[0].(models.UserKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteNetworkRequest indicates an expected call of ExecuteNetworkRequest
func (mr *MockNetworkDownStreamMockRecorder) ExecuteNetworkRequest(ctx, networkKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteNetworkRequest", reflect.TypeOf((*MockNetworkDownStream)(nil).ExecuteNetworkRequest), ctx, networkKey)
}
