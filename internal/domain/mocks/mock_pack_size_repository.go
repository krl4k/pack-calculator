// Code generated by MockGen. DO NOT EDIT.
// Source: calculate_product_packs/internal/domain (interfaces: PackSizeRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	domain "calculate_product_packs/internal/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPackSizeRepository is a mock of PackSizeRepository interface.
type MockPackSizeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPackSizeRepositoryMockRecorder
}

// MockPackSizeRepositoryMockRecorder is the mock recorder for MockPackSizeRepository.
type MockPackSizeRepositoryMockRecorder struct {
	mock *MockPackSizeRepository
}

// NewMockPackSizeRepository creates a new mock instance.
func NewMockPackSizeRepository(ctrl *gomock.Controller) *MockPackSizeRepository {
	mock := &MockPackSizeRepository{ctrl: ctrl}
	mock.recorder = &MockPackSizeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPackSizeRepository) EXPECT() *MockPackSizeRepositoryMockRecorder {
	return m.recorder
}

// GetPackSizes mocks base method.
func (m *MockPackSizeRepository) GetPackSizes() []domain.PackSize {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPackSizes")
	ret0, _ := ret[0].([]domain.PackSize)
	return ret0
}

// GetPackSizes indicates an expected call of GetPackSizes.
func (mr *MockPackSizeRepositoryMockRecorder) GetPackSizes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPackSizes", reflect.TypeOf((*MockPackSizeRepository)(nil).GetPackSizes))
}

// UpdatePackSizes mocks base method.
func (m *MockPackSizeRepository) UpdatePackSizes(arg0 []domain.PackSize) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePackSizes", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePackSizes indicates an expected call of UpdatePackSizes.
func (mr *MockPackSizeRepositoryMockRecorder) UpdatePackSizes(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePackSizes", reflect.TypeOf((*MockPackSizeRepository)(nil).UpdatePackSizes), arg0)
}
