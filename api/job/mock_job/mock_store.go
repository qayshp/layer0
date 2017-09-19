// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quintilesims/layer0/api/job (interfaces: Store)

// Package mock_job is a generated GoMock package.
package mock_job

import (
	gomock "github.com/golang/mock/gomock"
	job "github.com/quintilesims/layer0/api/job"
	models "github.com/quintilesims/layer0/common/models"
	reflect "reflect"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AcquireJob mocks base method
func (m *MockStore) AcquireJob(arg0 string) (bool, error) {
	ret := m.ctrl.Call(m, "AcquireJob", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcquireJob indicates an expected call of AcquireJob
func (mr *MockStoreMockRecorder) AcquireJob(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcquireJob", reflect.TypeOf((*MockStore)(nil).AcquireJob), arg0)
}

// Delete mocks base method
func (m *MockStore) Delete(arg0 string) error {
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockStoreMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStore)(nil).Delete), arg0)
}

// Insert mocks base method
func (m *MockStore) Insert(arg0 job.JobType, arg1 string) (string, error) {
	ret := m.ctrl.Call(m, "Insert", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert
func (mr *MockStoreMockRecorder) Insert(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockStore)(nil).Insert), arg0, arg1)
}

// SelectAll mocks base method
func (m *MockStore) SelectAll() ([]*models.Job, error) {
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll
func (mr *MockStoreMockRecorder) SelectAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockStore)(nil).SelectAll))
}

// SelectByID mocks base method
func (m *MockStore) SelectByID(arg0 string) (*models.Job, error) {
	ret := m.ctrl.Call(m, "SelectByID", arg0)
	ret0, _ := ret[0].(*models.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectByID indicates an expected call of SelectByID
func (mr *MockStoreMockRecorder) SelectByID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectByID", reflect.TypeOf((*MockStore)(nil).SelectByID), arg0)
}

// SetInsertHook mocks base method
func (m *MockStore) SetInsertHook(arg0 func(string)) {
	m.ctrl.Call(m, "SetInsertHook", arg0)
}

// SetInsertHook indicates an expected call of SetInsertHook
func (mr *MockStoreMockRecorder) SetInsertHook(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInsertHook", reflect.TypeOf((*MockStore)(nil).SetInsertHook), arg0)
}

// SetJobError mocks base method
func (m *MockStore) SetJobError(arg0 string, arg1 error) error {
	ret := m.ctrl.Call(m, "SetJobError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetJobError indicates an expected call of SetJobError
func (mr *MockStoreMockRecorder) SetJobError(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJobError", reflect.TypeOf((*MockStore)(nil).SetJobError), arg0, arg1)
}

// SetJobResult mocks base method
func (m *MockStore) SetJobResult(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "SetJobResult", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetJobResult indicates an expected call of SetJobResult
func (mr *MockStoreMockRecorder) SetJobResult(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJobResult", reflect.TypeOf((*MockStore)(nil).SetJobResult), arg0, arg1)
}

// SetJobStatus mocks base method
func (m *MockStore) SetJobStatus(arg0 string, arg1 job.Status) error {
	ret := m.ctrl.Call(m, "SetJobStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetJobStatus indicates an expected call of SetJobStatus
func (mr *MockStoreMockRecorder) SetJobStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetJobStatus", reflect.TypeOf((*MockStore)(nil).SetJobStatus), arg0, arg1)
}