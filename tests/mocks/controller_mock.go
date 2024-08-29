package mocks

import (
	reflect "reflect"

	fiber "github.com/gofiber/fiber/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockControllerInterface is a mock of ControllerInterface interface.
type MockControllerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockControllerInterfaceMockRecorder
}

// MockControllerInterfaceMockRecorder is the mock recorder for MockControllerInterface.
type MockControllerInterfaceMockRecorder struct {
	mock *MockControllerInterface
}

// NewMockControllerInterface creates a new mock instance.
func NewMockControllerInterface(ctrl *gomock.Controller) *MockControllerInterface {
	mock := &MockControllerInterface{ctrl: ctrl}
	mock.recorder = &MockControllerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockControllerInterface) EXPECT() *MockControllerInterfaceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockControllerInterface) CreateUser(ctx *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockControllerInterfaceMockRecorder) CreateUser(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockControllerInterface)(nil).CreateUser), ctx)
}

// LoginUser mocks base method.
func (m *MockControllerInterface) LoginUser(ctx *fiber.Ctx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockControllerInterfaceMockRecorder) LoginUser(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockControllerInterface)(nil).LoginUser), ctx)
}
