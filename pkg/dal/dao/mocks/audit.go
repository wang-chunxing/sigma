// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-sigma/sigma/pkg/dal/dao (interfaces: AuditService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/go-sigma/sigma/pkg/dal/models"
	gomock "go.uber.org/mock/gomock"
)

// MockAuditService is a mock of AuditService interface.
type MockAuditService struct {
	ctrl     *gomock.Controller
	recorder *MockAuditServiceMockRecorder
}

// MockAuditServiceMockRecorder is the mock recorder for MockAuditService.
type MockAuditServiceMockRecorder struct {
	mock *MockAuditService
}

// NewMockAuditService creates a new mock instance.
func NewMockAuditService(ctrl *gomock.Controller) *MockAuditService {
	mock := &MockAuditService{ctrl: ctrl}
	mock.recorder = &MockAuditServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuditService) EXPECT() *MockAuditServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuditService) Create(arg0 context.Context, arg1 *models.Audit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAuditServiceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuditService)(nil).Create), arg0, arg1)
}

// HotNamespace mocks base method.
func (m *MockAuditService) HotNamespace(arg0 context.Context, arg1 int64, arg2 int) ([]*models.Namespace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HotNamespace", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*models.Namespace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HotNamespace indicates an expected call of HotNamespace.
func (mr *MockAuditServiceMockRecorder) HotNamespace(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HotNamespace", reflect.TypeOf((*MockAuditService)(nil).HotNamespace), arg0, arg1, arg2)
}