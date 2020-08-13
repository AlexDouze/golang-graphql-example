// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics (interfaces: Client)

// Package mocks is a generated GoMock package.
package mocks

import (
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	http "net/http"
	reflect "reflect"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetPrometheusHTTPHandler mocks base method.
func (m *MockClient) GetPrometheusHTTPHandler() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrometheusHTTPHandler")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// GetPrometheusHTTPHandler indicates an expected call of GetPrometheusHTTPHandler.
func (mr *MockClientMockRecorder) GetPrometheusHTTPHandler() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrometheusHTTPHandler", reflect.TypeOf((*MockClient)(nil).GetPrometheusHTTPHandler))
}

// Instrument mocks base method.
func (m *MockClient) Instrument(arg0 string) gin.HandlerFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Instrument", arg0)
	ret0, _ := ret[0].(gin.HandlerFunc)
	return ret0
}

// Instrument indicates an expected call of Instrument.
func (mr *MockClientMockRecorder) Instrument(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Instrument", reflect.TypeOf((*MockClient)(nil).Instrument), arg0)
}
