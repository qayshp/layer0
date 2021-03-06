// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quintilesims/layer0/api/scheduler (interfaces: EnvironmentScaler)

// Package mock_scheduler is a generated GoMock package.
package mock_scheduler

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	models "github.com/quintilesims/layer0/common/models"
)

// MockEnvironmentScaler is a mock of EnvironmentScaler interface
type MockEnvironmentScaler struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentScalerMockRecorder
}

// MockEnvironmentScalerMockRecorder is the mock recorder for MockEnvironmentScaler
type MockEnvironmentScalerMockRecorder struct {
	mock *MockEnvironmentScaler
}

// NewMockEnvironmentScaler creates a new mock instance
func NewMockEnvironmentScaler(ctrl *gomock.Controller) *MockEnvironmentScaler {
	mock := &MockEnvironmentScaler{ctrl: ctrl}
	mock.recorder = &MockEnvironmentScalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentScaler) EXPECT() *MockEnvironmentScalerMockRecorder {
	return m.recorder
}

// Scale mocks base method
func (m *MockEnvironmentScaler) Scale(arg0 string) (*models.ScalerRunInfo, error) {
	ret := m.ctrl.Call(m, "Scale", arg0)
	ret0, _ := ret[0].(*models.ScalerRunInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scale indicates an expected call of Scale
func (mr *MockEnvironmentScalerMockRecorder) Scale(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockEnvironmentScaler)(nil).Scale), arg0)
}

// ScheduleRun mocks base method
func (m *MockEnvironmentScaler) ScheduleRun(arg0 string, arg1 time.Duration) {
	m.ctrl.Call(m, "ScheduleRun", arg0, arg1)
}

// ScheduleRun indicates an expected call of ScheduleRun
func (mr *MockEnvironmentScalerMockRecorder) ScheduleRun(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleRun", reflect.TypeOf((*MockEnvironmentScaler)(nil).ScheduleRun), arg0, arg1)
}
