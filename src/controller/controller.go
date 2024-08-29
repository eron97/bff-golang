package controller

import (
	"time"

	"github.com/eron97/bff-golang.git/cmd/config"
	"github.com/eron97/bff-golang.git/cmd/config/exceptions"
	"github.com/eron97/bff-golang.git/src/controller/dtos"
	"github.com/eron97/bff-golang.git/src/model/service"
	"github.com/eron97/bff-golang.git/src/view"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

func NewControllerInstance(serviceInterface service.ServiceInterface) ControllerInterface {
	return &Controller{
		service: serviceInterface,
	}
}

type ControllerInterface interface {
	CreateUser(ctx *fiber.Ctx) error
	LoginUser(ctx *fiber.Ctx) error
	RequestOtherService(ctx *fiber.Ctx) error
}

type Controller struct {
	service service.ServiceInterface
}

func (ctl *Controller) CreateUser(ctx *fiber.Ctx) error {

	createUser := ctx.Locals("createUser").(dtos.CreateUser)

	resp, err := ctl.service.CreateUserService(createUser)
	if err != nil {
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": "Error creating user: " + err.Error(),
		})
	}

	ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"usuário": view.ConvertDomainToResponse(resp),
	})

	return nil
}

func (ctl *Controller) LoginUser(ctx *fiber.Ctx) error {

	zap.L().Info("Starting user login")

	var user dtos.UserLogin

	if err := ctx.BodyParser(&user); err != nil {
		zap.L().Error("Error reading request data", zap.Error(err))
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unable to read request data",
		})
	}

	_, err := ctl.service.LoginUserService(user)
	if err != nil {
		zap.L().Error("Error when logging in", zap.Error(err))
		return ctx.Status(err.Code).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := GenerateToken(user)
	if err != nil {
		zap.L().Error("Error generating token", zap.Error(err))
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	zap.L().Info("Login successfully")
	ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successfully",
		"token":   token,
	})

	return nil
}

func (ctl *Controller) RequestOtherService(ctx *fiber.Ctx) error {

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Service found successfully",
	})
}

func GenerateToken(user dtos.UserLogin) (string, *exceptions.RestErr) {

	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"role":  "user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(config.PrivateKey)
	if err != nil {
		zap.L().Error("Error signing token", zap.Error(err))
		return "", exceptions.NewInternalServerError("Internal server error")
	}

	return tokenString, nil
}
