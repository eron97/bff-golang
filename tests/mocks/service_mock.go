package mocks

import (
	reflect "reflect"

	"github.com/eron97/bff-golang.git/cmd/config/exceptions"
	"github.com/eron97/bff-golang.git/src/controller/dtos"
	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	gomock "go.uber.org/mock/gomock"
)

// MockServiceInterface is a mock of ServiceInterface interface.
type MockServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInterfaceMockRecorder
}

// MockServiceInterfaceMockRecorder is the mock recorder for MockServiceInterface.
type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

// NewMockServiceInterface creates a new mock instance.
func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

// CreateUserService mocks base method.
func (m *MockServiceInterface) CreateUserService(request dtos.CreateUser) (*entity.CreateUser, *exceptions.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserService", request)
	ret0, _ := ret[0].(*entity.CreateUser)
	ret1, _ := ret[1].(*exceptions.RestErr)
	return ret0, ret1
}

// CreateUserService indicates an expected call of CreateUserService.
func (mr *MockServiceInterfaceMockRecorder) CreateUserService(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserService", reflect.TypeOf((*MockServiceInterface)(nil).CreateUserService), request)
}

// LoginUserService mocks base method.
func (m *MockServiceInterface) LoginUserService(request dtos.UserLogin) (bool, *exceptions.RestErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUserService", request)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*exceptions.RestErr)
	return ret0, ret1
}

// LoginUserService indicates an expected call of LoginUserService.
func (mr *MockServiceInterfaceMockRecorder) LoginUserService(request any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUserService", reflect.TypeOf((*MockServiceInterface)(nil).LoginUserService), request)
}
