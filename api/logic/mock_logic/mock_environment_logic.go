// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quintilesims/layer0/api/logic (interfaces: EnvironmentLogic)

// Package mock_logic is a generated GoMock package.
package mock_logic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/quintilesims/layer0/common/models"
)

// MockEnvironmentLogic is a mock of EnvironmentLogic interface
type MockEnvironmentLogic struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentLogicMockRecorder
}

// MockEnvironmentLogicMockRecorder is the mock recorder for MockEnvironmentLogic
type MockEnvironmentLogicMockRecorder struct {
	mock *MockEnvironmentLogic
}

// NewMockEnvironmentLogic creates a new mock instance
func NewMockEnvironmentLogic(ctrl *gomock.Controller) *MockEnvironmentLogic {
	mock := &MockEnvironmentLogic{ctrl: ctrl}
	mock.recorder = &MockEnvironmentLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentLogic) EXPECT() *MockEnvironmentLogicMockRecorder {
	return m.recorder
}

// CanCreateEnvironment mocks base method
func (m *MockEnvironmentLogic) CanCreateEnvironment(arg0 models.CreateEnvironmentRequest) (bool, error) {
	ret := m.ctrl.Call(m, "CanCreateEnvironment", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CanCreateEnvironment indicates an expected call of CanCreateEnvironment
func (mr *MockEnvironmentLogicMockRecorder) CanCreateEnvironment(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanCreateEnvironment", reflect.TypeOf((*MockEnvironmentLogic)(nil).CanCreateEnvironment), arg0)
}

// CreateEnvironment mocks base method
func (m *MockEnvironmentLogic) CreateEnvironment(arg0 models.CreateEnvironmentRequest) (*models.Environment, error) {
	ret := m.ctrl.Call(m, "CreateEnvironment", arg0)
	ret0, _ := ret[0].(*models.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnvironment indicates an expected call of CreateEnvironment
func (mr *MockEnvironmentLogicMockRecorder) CreateEnvironment(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockEnvironmentLogic)(nil).CreateEnvironment), arg0)
}

// CreateEnvironmentLink mocks base method
func (m *MockEnvironmentLogic) CreateEnvironmentLink(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "CreateEnvironmentLink", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEnvironmentLink indicates an expected call of CreateEnvironmentLink
func (mr *MockEnvironmentLogicMockRecorder) CreateEnvironmentLink(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironmentLink", reflect.TypeOf((*MockEnvironmentLogic)(nil).CreateEnvironmentLink), arg0, arg1)
}

// DeleteEnvironment mocks base method
func (m *MockEnvironmentLogic) DeleteEnvironment(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteEnvironment", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvironment indicates an expected call of DeleteEnvironment
func (mr *MockEnvironmentLogicMockRecorder) DeleteEnvironment(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironment", reflect.TypeOf((*MockEnvironmentLogic)(nil).DeleteEnvironment), arg0)
}

// DeleteEnvironmentLink mocks base method
func (m *MockEnvironmentLogic) DeleteEnvironmentLink(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "DeleteEnvironmentLink", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteEnvironmentLink indicates an expected call of DeleteEnvironmentLink
func (mr *MockEnvironmentLogicMockRecorder) DeleteEnvironmentLink(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEnvironmentLink", reflect.TypeOf((*MockEnvironmentLogic)(nil).DeleteEnvironmentLink), arg0, arg1)
}

// GetEnvironment mocks base method
func (m *MockEnvironmentLogic) GetEnvironment(arg0 string) (*models.Environment, error) {
	ret := m.ctrl.Call(m, "GetEnvironment", arg0)
	ret0, _ := ret[0].(*models.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment
func (mr *MockEnvironmentLogicMockRecorder) GetEnvironment(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockEnvironmentLogic)(nil).GetEnvironment), arg0)
}

// ListEnvironments mocks base method
func (m *MockEnvironmentLogic) ListEnvironments() ([]models.EnvironmentSummary, error) {
	ret := m.ctrl.Call(m, "ListEnvironments")
	ret0, _ := ret[0].([]models.EnvironmentSummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironments indicates an expected call of ListEnvironments
func (mr *MockEnvironmentLogicMockRecorder) ListEnvironments() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironments", reflect.TypeOf((*MockEnvironmentLogic)(nil).ListEnvironments))
}

// UpdateEnvironment mocks base method
func (m *MockEnvironmentLogic) UpdateEnvironment(arg0 string, arg1 int) (*models.Environment, error) {
	ret := m.ctrl.Call(m, "UpdateEnvironment", arg0, arg1)
	ret0, _ := ret[0].(*models.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEnvironment indicates an expected call of UpdateEnvironment
func (mr *MockEnvironmentLogicMockRecorder) UpdateEnvironment(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvironment", reflect.TypeOf((*MockEnvironmentLogic)(nil).UpdateEnvironment), arg0, arg1)
}
