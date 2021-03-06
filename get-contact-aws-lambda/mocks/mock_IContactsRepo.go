// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/wilomendez/Onboarding/get-contact-aws-lambda/models (interfaces: IContactsRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/wilomendez/Onboarding/get-contact-aws-lambda/models"
)

// MockIContactsRepo is a mock of IContactsRepo interface.
type MockIContactsRepo struct {
	ctrl     *gomock.Controller
	recorder *MockIContactsRepoMockRecorder
}

// MockIContactsRepoMockRecorder is the mock recorder for MockIContactsRepo.
type MockIContactsRepoMockRecorder struct {
	mock *MockIContactsRepo
}

// NewMockIContactsRepo creates a new mock instance.
func NewMockIContactsRepo(ctrl *gomock.Controller) *MockIContactsRepo {
	mock := &MockIContactsRepo{ctrl: ctrl}
	mock.recorder = &MockIContactsRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIContactsRepo) EXPECT() *MockIContactsRepoMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockIContactsRepo) Find(arg0 string) (models.Contacts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].(models.Contacts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockIContactsRepoMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIContactsRepo)(nil).Find), arg0)
}
