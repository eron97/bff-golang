package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/eron97/bff-golang.git/cmd/config/exceptions"
	"github.com/eron97/bff-golang.git/src/controller"
	"github.com/eron97/bff-golang.git/src/controller/dtos"
	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	"github.com/eron97/bff-golang.git/tests/mocks"

	"github.com/go-playground/assert"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/mock/gomock"
)

func TestCreateUserSuccess(t *testing.T) {
	app := fiber.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	mockController := controller.NewControllerInstance(mockService)

	mockUser := dtos.CreateUser{
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		Password:   "password",
	}

	app.Post("/register", func(c *fiber.Ctx) error {
		// Definir o valor de "createUser" em ctx.Locals
		c.Locals("createUser", mockUser)
		return mockController.CreateUser(c)
	})

	// Mock do serviço para simular sucesso ao criar o usuário
	mockService.EXPECT().CreateUserService(mockUser).Return(&entity.CreateUser{}, nil)

	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(`{"First_Name":"John",
	"Last_Name":"Doe","Email":"john.doe@example.com",
	"Password":"password"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	// Verificar o status de retorno
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	// Verificar o corpo da resposta
	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)
	assert.Equal(t, "User created successfully", responseBody["message"])
}
func TestCreateUserError(t *testing.T) {
	app := fiber.New()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	mockController := controller.NewControllerInstance(mockService)

	app.Post("/register", func(c *fiber.Ctx) error {

		mockUser := dtos.CreateUser{
			First_Name: "John",
			Last_Name:  "Doe",
			Email:      "john.doe@example.com",
			Password:   "password",
		}
		c.Locals("createUser", mockUser)
		return mockController.CreateUser(c)
	})

	mockUser := dtos.CreateUser{
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		Password:   "password",
	}

	mockService.EXPECT().CreateUserService(mockUser).Return(nil, exceptions.NewBadRequestError("Email is already associated with an existing account"))

	req := httptest.NewRequest("POST", "/register", bytes.NewBufferString(`{"First_Name":"John","Last_Name":"Doe",
	"Email":"john.doe@example.com","Password":"password"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)

	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expectedErrorMessage := "Error creating user: Email is already associated with an existing account"
	assert.Equal(t, expectedErrorMessage, responseBody["error"])
}

func TestLoginUserSuccess(t *testing.T) {
	app := fiber.New()

	// Criar o controlador de mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockServiceInterface(ctrl)

	// Criar o mock do controlador
	mockController := controller.NewControllerInstance(mockService)

	app.Post("/login", func(c *fiber.Ctx) error {
		c.Locals("ip", "127.0.0.1") // Configurar um IP mockado
		return mockController.LoginUser(c)
	})

	mockUser := dtos.UserLogin{
		Email:    "john.doe@example.com",
		Password: "password",
	}

	// Simular um sucesso no serviço de login
	mockService.EXPECT().LoginUserService(mockUser).Return(true, nil)

	req := httptest.NewRequest("POST", "/login", bytes.NewBufferString(`{"email":"john.doe@example.com","password":"password"}`))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req)

	// Verificar o status de retorno
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	// Deserializar o corpo da resposta
	var responseBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&responseBody)

	// Verificar a mensagem de sucesso
	assert.Equal(t, "Login successfully", responseBody["message"])
}
