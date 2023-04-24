// Code generated by MockGen. DO NOT EDIT.
// Source: ../../internal/repositories/openaq/interface.go

// Package openaqmocks is a generated GoMock package.
package openaqmocks

import (
	entities "TPBDM/scraper/internal/entities"
	openaq "TPBDM/scraper/internal/repositories/openaq"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// GetMeasurements mocks base method.
func (m *MockRepository) GetMeasurements(query openaq.QueryContract) ([]entities.Measurement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeasurements", query)
	ret0, _ := ret[0].([]entities.Measurement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeasurements indicates an expected call of GetMeasurements.
func (mr *MockRepositoryMockRecorder) GetMeasurements(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeasurements", reflect.TypeOf((*MockRepository)(nil).GetMeasurements), query)
}

// GetMeasurementsCount mocks base method.
func (m *MockRepository) GetMeasurementsCount(query openaq.QueryContract) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeasurementsCount", query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeasurementsCount indicates an expected call of GetMeasurementsCount.
func (mr *MockRepositoryMockRecorder) GetMeasurementsCount(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeasurementsCount", reflect.TypeOf((*MockRepository)(nil).GetMeasurementsCount), query)
}
