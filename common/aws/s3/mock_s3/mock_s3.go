// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quintilesims/layer0/common/aws/s3 (interfaces: Provider)

// Package mock_s3 is a generated GoMock package.
package mock_s3

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockProvider is a mock of Provider interface
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// DeleteObject mocks base method
func (m *MockProvider) DeleteObject(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "DeleteObject", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObject indicates an expected call of DeleteObject
func (mr *MockProviderMockRecorder) DeleteObject(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockProvider)(nil).DeleteObject), arg0, arg1)
}

// GetObject mocks base method
func (m *MockProvider) GetObject(arg0, arg1 string) ([]byte, error) {
	ret := m.ctrl.Call(m, "GetObject", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject
func (mr *MockProviderMockRecorder) GetObject(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockProvider)(nil).GetObject), arg0, arg1)
}

// GetObjectToFile mocks base method
func (m *MockProvider) GetObjectToFile(arg0, arg1, arg2 string, arg3 os.FileMode) error {
	ret := m.ctrl.Call(m, "GetObjectToFile", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetObjectToFile indicates an expected call of GetObjectToFile
func (mr *MockProviderMockRecorder) GetObjectToFile(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObjectToFile", reflect.TypeOf((*MockProvider)(nil).GetObjectToFile), arg0, arg1, arg2, arg3)
}

// ListObjects mocks base method
func (m *MockProvider) ListObjects(arg0, arg1 string) ([]string, error) {
	ret := m.ctrl.Call(m, "ListObjects", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListObjects indicates an expected call of ListObjects
func (mr *MockProviderMockRecorder) ListObjects(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListObjects", reflect.TypeOf((*MockProvider)(nil).ListObjects), arg0, arg1)
}

// PutObject mocks base method
func (m *MockProvider) PutObject(arg0, arg1 string, arg2 []byte) error {
	ret := m.ctrl.Call(m, "PutObject", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObject indicates an expected call of PutObject
func (mr *MockProviderMockRecorder) PutObject(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockProvider)(nil).PutObject), arg0, arg1, arg2)
}

// PutObjectFromFile mocks base method
func (m *MockProvider) PutObjectFromFile(arg0, arg1, arg2 string) error {
	ret := m.ctrl.Call(m, "PutObjectFromFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutObjectFromFile indicates an expected call of PutObjectFromFile
func (mr *MockProviderMockRecorder) PutObjectFromFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObjectFromFile", reflect.TypeOf((*MockProvider)(nil).PutObjectFromFile), arg0, arg1, arg2)
}