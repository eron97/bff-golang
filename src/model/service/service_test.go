package service_test

import (
	"errors"
	"testing"

	"github.com/eron97/bff-golang.git/cmd/config/exceptions"
	"github.com/eron97/bff-golang.git/src/controller/dtos"
	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	"github.com/eron97/bff-golang.git/src/model/service"
	"github.com/eron97/bff-golang.git/tests/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mocks.NewMockPersistenceInterface(ctrl)

	cryptoService := &MockCrypto{}
	srv := service.NewServiceInstance(cryptoService, repo)

	user := dtos.CreateUser{
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		CepBR:      "12345678",
		Country:    "Brazil",
		City:       "São Paulo",
		Address:    "Rua das Flores, 123",
		Password:   "password",
	}

	hash, _ := cryptoService.HashPassword(user.Password)

	expectedUser := &entity.CreateUser{
		First_Name: user.First_Name,
		Last_Name:  user.Last_Name,
		Email:      user.Email,
		CepBR:      user.CepBR,
		Country:    user.Country,
		City:       user.City,
		Address:    user.Address,
		Password:   hash,
	}

	repo.EXPECT().GetUser(user.Email).Return(&entity.CreateUser{Email: ""})
	repo.EXPECT().CreateUser(gomock.Any()).Return(nil)

	resp, err := srv.CreateUserService(user)

	assert.Nil(t, err, "unexpected error was returned")
	assert.Equal(t, expectedUser.Email, resp.Email)
}

func TestCreateUserService_EmailExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cryptoService := &MockCrypto{}

	repo := mocks.NewMockPersistenceInterface(ctrl)
	srv := service.NewServiceInstance(cryptoService, repo)

	user := dtos.CreateUser{
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		CepBR:      "12345678",
		Country:    "Brazil",
		City:       "São Paulo",
		Address:    "Rua das Flores, 123",
		Password:   "password",
	}

	repo.EXPECT().GetUser(user.Email).Return(&entity.CreateUser{Email: user.Email})

	_, err := srv.CreateUserService(user)

	assert.Equal(t, "Email is already associated with an existing account", err.Message)
}

func TestCreateUserService_HashPasswordError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockPersistenceInterface(ctrl)

	cryptoService := &MockCrypto{ShouldFail: true}

	srv := service.NewServiceInstance(cryptoService, repo)

	user := dtos.CreateUser{
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		CepBR:      "12345678",
		Country:    "Brazil",
		City:       "São Paulo",
		Address:    "Rua das Flores, 123",
		Password:   "password",
	}

	repo.EXPECT().GetUser(user.Email).Return(&entity.CreateUser{Email: ""})

	_, err := srv.CreateUserService(user)

	assert.NotNil(t, err, "expected an error, but it was not returned")
	assert.Equal(t, "Internal server error", err.Message)
}

func TestCreateUserService_CreateUserError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockPersistenceInterface(ctrl)

	cryptoService := &MockCrypto{}

	srv := service.NewServiceInstance(cryptoService, repo)

	user := dtos.CreateUser{
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		CepBR:      "12345678",
		Country:    "Brazil",
		City:       "São Paulo",
		Address:    "Rua das Flores, 123",
		Password:   "password",
	}

	repo.EXPECT().GetUser(user.Email).Return(&entity.CreateUser{Email: ""})

	expectedError := errors.New("Internal server error")

	repo.EXPECT().CreateUser(gomock.Any()).Return(expectedError)

	resp, err := srv.CreateUserService(user)
	assert.Nil(t, resp, "expected zero as an answer")
	assert.NotNil(t, err, "expected an error, but it was not returned")
	assert.IsType(t, &exceptions.RestErr{}, err)
}

func TestLoginUserService_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockPersistenceInterface(ctrl)

	cryptoService := &MockCrypto{}

	srv := service.NewServiceInstance(cryptoService, repo)

	user := dtos.UserLogin{
		Email:    "john.doe@example.com",
		Password: "password",
	}

	hashedPassword, err := cryptoService.HashPassword(user.Email)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	mockUser := &entity.CreateUser{
		Email:    user.Email,
		Password: hashedPassword,
	}

	repo.EXPECT().GetUser(user.Email).Return(mockUser)

	resp, err := srv.LoginUserService(user)

	assert.Equal(t, true, resp)
	assert.Nil(t, err, "expected an error, but it was not returned: %v", err)
}

func TestLoginUserService_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockPersistenceInterface(ctrl)

	cryptoService := &MockCrypto{}

	srv := service.NewServiceInstance(cryptoService, repo)

	user := dtos.UserLogin{
		Email:    "john.doe@example.com",
		Password: "false",
	}

	repo.EXPECT().GetUser(user.Email).Return(&entity.CreateUser{Email: ""})

	resp, err := srv.LoginUserService(user)

	assert.Equal(t, false, resp)
	assert.Equal(t, "Account not found", err.Message)
}

func TestLoginUserService_IncorrectPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockPersistenceInterface(ctrl)

	cryptoService := &MockCrypto{ShouldFail: true}

	srv := service.NewServiceInstance(cryptoService, repo)

	user := dtos.UserLogin{
		Email:    "john.doe@example.com",
		Password: "false",
	}

	repo.EXPECT().GetUser(user.Email).Return(&entity.CreateUser{Email: user.Email})

	cryptoService.CheckPassword(user.Password, "hashedPassword123")

	resp, err := srv.LoginUserService(user)

	assert.Equal(t, false, resp)
	assert.Equal(t, "The password entered is incorrect", err.Message)
}

type MockCrypto struct {
	ShouldFail bool
}

func (m *MockCrypto) HashPassword(password string) (string, error) {
	if m.ShouldFail {
		return "", errors.New("error creating password hash")
	}
	return "hashedPassword", nil
}

func (m *MockCrypto) CheckPassword(plainPassword, hashedPassword string) (bool, error) {
	if m.ShouldFail {
		return false, errors.New("The password entered is incorrect")
	}
	return true, nil
}
