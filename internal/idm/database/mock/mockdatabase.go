// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/borisbbtest/GoMon/internal/idm/database (interfaces: Storager)

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	configs "github.com/borisbbtest/GoMon/internal/idm/configs"
	idm "github.com/borisbbtest/GoMon/internal/models/idm"
	gomock "github.com/golang/mock/gomock"
)

// MockStorager is a mock of Storager interface.
type MockStorager struct {
	ctrl     *gomock.Controller
	recorder *MockStoragerMockRecorder
}

// MockStoragerMockRecorder is the mock recorder for MockStorager.
type MockStoragerMockRecorder struct {
	mock *MockStorager
}

// NewMockStorager creates a new mock instance.
func NewMockStorager(ctrl *gomock.Controller) *MockStorager {
	mock := &MockStorager{ctrl: ctrl}
	mock.recorder = &MockStoragerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorager) EXPECT() *MockStoragerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStorager) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockStoragerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStorager)(nil).Close))
}

// CreateSession mocks base method.
func (m *MockStorager) CreateSession(arg0 context.Context, arg1 *configs.AppConfig, arg2 *idm.Session) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockStoragerMockRecorder) CreateSession(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockStorager)(nil).CreateSession), arg0, arg1, arg2)
}

// CreateTables mocks base method.
func (m *MockStorager) CreateTables(arg0 context.Context, arg1 *configs.AppConfig) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTables", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateTables indicates an expected call of CreateTables.
func (mr *MockStoragerMockRecorder) CreateTables(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTables", reflect.TypeOf((*MockStorager)(nil).CreateTables), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStorager) CreateUser(arg0 context.Context, arg1 *configs.AppConfig, arg2 *idm.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoragerMockRecorder) CreateUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorager)(nil).CreateUser), arg0, arg1, arg2)
}

// DeleteSession mocks base method.
func (m *MockStorager) DeleteSession(arg0 context.Context, arg1 *configs.AppConfig, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockStoragerMockRecorder) DeleteSession(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockStorager)(nil).DeleteSession), arg0, arg1, arg2, arg3)
}

// DeleteUser mocks base method.
func (m *MockStorager) DeleteUser(arg0 context.Context, arg1 *configs.AppConfig, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoragerMockRecorder) DeleteUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStorager)(nil).DeleteUser), arg0, arg1, arg2)
}

// GetAllSessions mocks base method.
func (m *MockStorager) GetAllSessions(arg0 context.Context, arg1 *configs.AppConfig) ([]*idm.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSessions", arg0, arg1)
	ret0, _ := ret[0].([]*idm.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSessions indicates an expected call of GetAllSessions.
func (mr *MockStoragerMockRecorder) GetAllSessions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSessions", reflect.TypeOf((*MockStorager)(nil).GetAllSessions), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockStorager) GetAllUsers(arg0 context.Context, arg1 *configs.AppConfig) ([]*idm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0, arg1)
	ret0, _ := ret[0].([]*idm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockStoragerMockRecorder) GetAllUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockStorager)(nil).GetAllUsers), arg0, arg1)
}

// GetSession mocks base method.
func (m *MockStorager) GetSession(arg0 context.Context, arg1 *configs.AppConfig, arg2, arg3 string) (*idm.Session, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSession", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*idm.Session)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSession indicates an expected call of GetSession.
func (mr *MockStoragerMockRecorder) GetSession(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSession", reflect.TypeOf((*MockStorager)(nil).GetSession), arg0, arg1, arg2, arg3)
}

// GetUser mocks base method.
func (m *MockStorager) GetUser(arg0 context.Context, arg1 *configs.AppConfig, arg2 string) (*idm.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(*idm.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoragerMockRecorder) GetUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStorager)(nil).GetUser), arg0, arg1, arg2)
}
