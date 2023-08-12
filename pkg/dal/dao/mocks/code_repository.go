// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/go-sigma/sigma/pkg/dal/dao (interfaces: CodeRepositoryService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	models "github.com/go-sigma/sigma/pkg/dal/models"
	types "github.com/go-sigma/sigma/pkg/types"
	enums "github.com/go-sigma/sigma/pkg/types/enums"
	gomock "go.uber.org/mock/gomock"
)

// MockCodeRepositoryService is a mock of CodeRepositoryService interface.
type MockCodeRepositoryService struct {
	ctrl     *gomock.Controller
	recorder *MockCodeRepositoryServiceMockRecorder
}

// MockCodeRepositoryServiceMockRecorder is the mock recorder for MockCodeRepositoryService.
type MockCodeRepositoryServiceMockRecorder struct {
	mock *MockCodeRepositoryService
}

// NewMockCodeRepositoryService creates a new mock instance.
func NewMockCodeRepositoryService(ctrl *gomock.Controller) *MockCodeRepositoryService {
	mock := &MockCodeRepositoryService{ctrl: ctrl}
	mock.recorder = &MockCodeRepositoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCodeRepositoryService) EXPECT() *MockCodeRepositoryServiceMockRecorder {
	return m.recorder
}

// CreateInBatches mocks base method.
func (m *MockCodeRepositoryService) CreateInBatches(arg0 context.Context, arg1 []*models.CodeRepository) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInBatches", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInBatches indicates an expected call of CreateInBatches.
func (mr *MockCodeRepositoryServiceMockRecorder) CreateInBatches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInBatches", reflect.TypeOf((*MockCodeRepositoryService)(nil).CreateInBatches), arg0, arg1)
}

// CreateOwnersInBatches mocks base method.
func (m *MockCodeRepositoryService) CreateOwnersInBatches(arg0 context.Context, arg1 []*models.CodeRepositoryOwner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOwnersInBatches", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOwnersInBatches indicates an expected call of CreateOwnersInBatches.
func (mr *MockCodeRepositoryServiceMockRecorder) CreateOwnersInBatches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOwnersInBatches", reflect.TypeOf((*MockCodeRepositoryService)(nil).CreateOwnersInBatches), arg0, arg1)
}

// DeleteInBatches mocks base method.
func (m *MockCodeRepositoryService) DeleteInBatches(arg0 context.Context, arg1 []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteInBatches", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteInBatches indicates an expected call of DeleteInBatches.
func (mr *MockCodeRepositoryServiceMockRecorder) DeleteInBatches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteInBatches", reflect.TypeOf((*MockCodeRepositoryService)(nil).DeleteInBatches), arg0, arg1)
}

// DeleteOwnerInBatches mocks base method.
func (m *MockCodeRepositoryService) DeleteOwnerInBatches(arg0 context.Context, arg1 []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOwnerInBatches", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOwnerInBatches indicates an expected call of DeleteOwnerInBatches.
func (mr *MockCodeRepositoryServiceMockRecorder) DeleteOwnerInBatches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOwnerInBatches", reflect.TypeOf((*MockCodeRepositoryService)(nil).DeleteOwnerInBatches), arg0, arg1)
}

// ListAll mocks base method.
func (m *MockCodeRepositoryService) ListAll(arg0 context.Context, arg1 int64) ([]*models.CodeRepository, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0, arg1)
	ret0, _ := ret[0].([]*models.CodeRepository)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll.
func (mr *MockCodeRepositoryServiceMockRecorder) ListAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockCodeRepositoryService)(nil).ListAll), arg0, arg1)
}

// ListOwnerWithPagination mocks base method.
func (m *MockCodeRepositoryService) ListOwnerWithPagination(arg0 context.Context, arg1 int64, arg2 enums.Provider, arg3 *string, arg4 types.Pagination, arg5 types.Sortable) ([]*models.CodeRepositoryOwner, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOwnerWithPagination", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*models.CodeRepositoryOwner)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListOwnerWithPagination indicates an expected call of ListOwnerWithPagination.
func (mr *MockCodeRepositoryServiceMockRecorder) ListOwnerWithPagination(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOwnerWithPagination", reflect.TypeOf((*MockCodeRepositoryService)(nil).ListOwnerWithPagination), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListOwnersAll mocks base method.
func (m *MockCodeRepositoryService) ListOwnersAll(arg0 context.Context, arg1 int64) ([]*models.CodeRepositoryOwner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOwnersAll", arg0, arg1)
	ret0, _ := ret[0].([]*models.CodeRepositoryOwner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOwnersAll indicates an expected call of ListOwnersAll.
func (mr *MockCodeRepositoryServiceMockRecorder) ListOwnersAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOwnersAll", reflect.TypeOf((*MockCodeRepositoryService)(nil).ListOwnersAll), arg0, arg1)
}

// ListWithPagination mocks base method.
func (m *MockCodeRepositoryService) ListWithPagination(arg0 context.Context, arg1 int64, arg2 enums.Provider, arg3, arg4 *string, arg5 types.Pagination, arg6 types.Sortable) ([]*models.CodeRepository, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWithPagination", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].([]*models.CodeRepository)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListWithPagination indicates an expected call of ListWithPagination.
func (mr *MockCodeRepositoryServiceMockRecorder) ListWithPagination(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWithPagination", reflect.TypeOf((*MockCodeRepositoryService)(nil).ListWithPagination), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}