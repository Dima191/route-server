// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/respository.go
//
// Generated by this command:
//
//	mockgen -source internal/repository/respository.go -destination internal/repository/respository_mock.go
//

// Package mock_repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"
	models "github.com/Dima191/route-server/internal/models"

	gomock "go.uber.org/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AllRoutes mocks base method.
func (m *MockRepository) AllRoutes(ctx context.Context) (chan models.Route, chan error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRoutes", ctx)
	ret0, _ := ret[0].(chan models.Route)
	ret1, _ := ret[1].(chan error)
	return ret0, ret1
}

// AllRoutes indicates an expected call of AllRoutes.
func (mr *MockRepositoryMockRecorder) AllRoutes(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRoutes", reflect.TypeOf((*MockRepository)(nil).AllRoutes), ctx)
}

// Close mocks base method.
func (m *MockRepository) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRepositoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepository)(nil).Close))
}