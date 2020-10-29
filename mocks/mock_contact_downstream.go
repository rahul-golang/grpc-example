// Code generated by MockGen. DO NOT EDIT.
// Source: downstream/contact_downstream.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	models "interview/models"
	reflect "reflect"
)

// MockContactDownStream is a mock of ContactDownStream interface
type MockContactDownStream struct {
	ctrl     *gomock.Controller
	recorder *MockContactDownStreamMockRecorder
}

// MockContactDownStreamMockRecorder is the mock recorder for MockContactDownStream
type MockContactDownStreamMockRecorder struct {
	mock *MockContactDownStream
}

// NewMockContactDownStream creates a new mock instance
func NewMockContactDownStream(ctrl *gomock.Controller) *MockContactDownStream {
	mock := &MockContactDownStream{ctrl: ctrl}
	mock.recorder = &MockContactDownStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockContactDownStream) EXPECT() *MockContactDownStreamMockRecorder {
	return m.recorder
}

// ExecuteContactRequest mocks base method
func (m *MockContactDownStream) ExecuteContactRequest(ctx context.Context, key1, key2 int64) (models.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteContactRequest", ctx, key1, key2)
	ret0, _ := ret[0].(models.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteContactRequest indicates an expected call of ExecuteContactRequest
func (mr *MockContactDownStreamMockRecorder) ExecuteContactRequest(ctx, key1, key2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteContactRequest", reflect.TypeOf((*MockContactDownStream)(nil).ExecuteContactRequest), ctx, key1, key2)
}
