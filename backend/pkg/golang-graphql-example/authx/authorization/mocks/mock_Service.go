// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/authx/authorization (interfaces: Service)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// CheckAuthorized mocks base method
func (m *MockService) CheckAuthorized(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckAuthorized", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckAuthorized indicates an expected call of CheckAuthorized
func (mr *MockServiceMockRecorder) CheckAuthorized(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckAuthorized", reflect.TypeOf((*MockService)(nil).CheckAuthorized), arg0, arg1, arg2)
}

// IsAuthorized mocks base method
func (m *MockService) IsAuthorized(arg0 context.Context, arg1, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsAuthorized", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsAuthorized indicates an expected call of IsAuthorized
func (mr *MockServiceMockRecorder) IsAuthorized(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsAuthorized", reflect.TypeOf((*MockService)(nil).IsAuthorized), arg0, arg1, arg2)
}

// Middleware mocks base method
func (m *MockService) Middleware() gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Middleware")
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Middleware indicates an expected call of Middleware
func (mr *MockServiceMockRecorder) Middleware() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Middleware", reflect.TypeOf((*MockService)(nil).Middleware))
}
