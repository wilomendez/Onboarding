// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wilomendez/Onboarding/get-contact-aws-lambda/controllers (interfaces: IGetContactHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/wilomendez/Onboarding/get-contact-aws-lambda/models"
)

// MockIGetContactHandler is a mock of IGetContactHandler interface.
type MockIGetContactHandler struct {
	ctrl     *gomock.Controller
	recorder *MockIGetContactHandlerMockRecorder
}

// MockIGetContactHandlerMockRecorder is the mock recorder for MockIGetContactHandler.
type MockIGetContactHandlerMockRecorder struct {
	mock *MockIGetContactHandler
}

// NewMockIGetContactHandler creates a new mock instance.
func NewMockIGetContactHandler(ctrl *gomock.Controller) *MockIGetContactHandler {
	mock := &MockIGetContactHandler{ctrl: ctrl}
	mock.recorder = &MockIGetContactHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGetContactHandler) EXPECT() *MockIGetContactHandlerMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockIGetContactHandler) Find(arg0 context.Context, arg1 models.Request) (models.Contacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0, arg1)
	ret0, _ := ret[0].(models.Contacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIGetContactHandlerMockRecorder) Find(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIGetContactHandler)(nil).Find), arg0, arg1)
}
