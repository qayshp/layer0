// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quintilesims/layer0/api/logic (interfaces: DeployLogic)

// Package mock_logic is a generated GoMock package.
package mock_logic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/quintilesims/layer0/common/models"
)

// MockDeployLogic is a mock of DeployLogic interface
type MockDeployLogic struct {
	ctrl     *gomock.Controller
	recorder *MockDeployLogicMockRecorder
}

// MockDeployLogicMockRecorder is the mock recorder for MockDeployLogic
type MockDeployLogicMockRecorder struct {
	mock *MockDeployLogic
}

// NewMockDeployLogic creates a new mock instance
func NewMockDeployLogic(ctrl *gomock.Controller) *MockDeployLogic {
	mock := &MockDeployLogic{ctrl: ctrl}
	mock.recorder = &MockDeployLogicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeployLogic) EXPECT() *MockDeployLogicMockRecorder {
	return m.recorder
}

// CreateDeploy mocks base method
func (m *MockDeployLogic) CreateDeploy(arg0 models.CreateDeployRequest) (*models.Deploy, error) {
	ret := m.ctrl.Call(m, "CreateDeploy", arg0)
	ret0, _ := ret[0].(*models.Deploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeploy indicates an expected call of CreateDeploy
func (mr *MockDeployLogicMockRecorder) CreateDeploy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeploy", reflect.TypeOf((*MockDeployLogic)(nil).CreateDeploy), arg0)
}

// DeleteDeploy mocks base method
func (m *MockDeployLogic) DeleteDeploy(arg0 string) error {
	ret := m.ctrl.Call(m, "DeleteDeploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeploy indicates an expected call of DeleteDeploy
func (mr *MockDeployLogicMockRecorder) DeleteDeploy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeploy", reflect.TypeOf((*MockDeployLogic)(nil).DeleteDeploy), arg0)
}

// GetDeploy mocks base method
func (m *MockDeployLogic) GetDeploy(arg0 string) (*models.Deploy, error) {
	ret := m.ctrl.Call(m, "GetDeploy", arg0)
	ret0, _ := ret[0].(*models.Deploy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploy indicates an expected call of GetDeploy
func (mr *MockDeployLogicMockRecorder) GetDeploy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploy", reflect.TypeOf((*MockDeployLogic)(nil).GetDeploy), arg0)
}

// ListDeploys mocks base method
func (m *MockDeployLogic) ListDeploys() ([]*models.DeploySummary, error) {
	ret := m.ctrl.Call(m, "ListDeploys")
	ret0, _ := ret[0].([]*models.DeploySummary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeploys indicates an expected call of ListDeploys
func (mr *MockDeployLogicMockRecorder) ListDeploys() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeploys", reflect.TypeOf((*MockDeployLogic)(nil).ListDeploys))
}