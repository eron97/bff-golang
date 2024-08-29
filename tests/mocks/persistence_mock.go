package mocks

import (
	reflect "reflect"

	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	gomock "go.uber.org/mock/gomock"
)

// MockPersistenceInterface is a mock of PersistenceInterface interface.
type MockPersistenceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPersistenceInterfaceMockRecorder
}

// MockPersistenceInterfaceMockRecorder is the mock recorder for MockPersistenceInterface.
type MockPersistenceInterfaceMockRecorder struct {
	mock *MockPersistenceInterface
}

// NewMockPersistenceInterface creates a new mock instance.
func NewMockPersistenceInterface(ctrl *gomock.Controller) *MockPersistenceInterface {
	mock := &MockPersistenceInterface{ctrl: ctrl}
	mock.recorder = &MockPersistenceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersistenceInterface) EXPECT() *MockPersistenceInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockPersistenceInterface) CreateUser(user entity.CreateUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockPersistenceInterfaceMockRecorder) CreateUser(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockPersistenceInterface)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockPersistenceInterface) GetUser(email string) *entity.CreateUser {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", email)
	ret0, _ := ret[0].(*entity.CreateUser)
	return ret0
}

// GetUser indicates an expected call of GetUser.
func (mr *MockPersistenceInterfaceMockRecorder) GetUser(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockPersistenceInterface)(nil).GetUser), email)
}

// VerifyExist mocks base method.
func (m *MockPersistenceInterface) VerifyExist(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyExist", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyExist indicates an expected call of VerifyExist.
func (mr *MockPersistenceInterfaceMockRecorder) VerifyExist(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyExist", reflect.TypeOf((*MockPersistenceInterface)(nil).VerifyExist), email)
}
