package service

import (
	"time"

	"github.com/eron97/bff-golang.git/cmd/config/exceptions"
	"github.com/eron97/bff-golang.git/src/controller/dtos"
	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	"github.com/eron97/bff-golang.git/src/model/persistence"
	"github.com/eron97/bff-golang.git/src/model/service/crypto"
	"go.uber.org/zap"
)

type ServiceInterface interface {
	CreateUserService(request dtos.CreateUser) (*entity.CreateUser, *exceptions.RestErr)
	LoginUserService(request dtos.UserLogin) (bool, *exceptions.RestErr)
}

type Service struct {
	crypto crypto.CryptoInterface
	db     persistence.PersistenceInterface
}

func NewServiceInstance(crypto crypto.CryptoInterface, db persistence.PersistenceInterface) ServiceInterface {
	return &Service{
		crypto: crypto,
		db:     db,
	}
}

func (srv *Service) LoginUserService(request dtos.UserLogin) (bool, *exceptions.RestErr) {
	zap.L().Info("Starting login service")

	user := srv.db.GetUser(request.Email)

	if user.Email == "" {
		zap.L().Warn("User not found", zap.String("email", request.Email))
		return false, exceptions.NewNotFoundError("Account not found")
	}

	_, err := srv.crypto.CheckPassword(request.Password, user.Password)
	if err != nil {
		zap.L().Warn("Incorrect password", zap.String("email", request.Email))
		return false, exceptions.NewUnauthorizedRequestError("The password entered is incorrect")
	}

	zap.L().Info("Successful login", zap.String("email", request.Email))
	return true, nil
}

func buildUserEntity(request dtos.CreateUser, hashedPassword string) *entity.CreateUser {
	return &entity.CreateUser{
		ID:         entity.NewID(),
		First_Name: request.First_Name,
		Last_Name:  request.Last_Name,
		Email:      request.Email,
		CepBR:      request.CepBR,
		Country:    request.Country,
		City:       request.City,
		Address:    request.Address,
		Password:   hashedPassword,
		CreateAt:   time.Now(),
	}
}

func (srv *Service) CreateUserService(request dtos.CreateUser) (*entity.CreateUser, *exceptions.RestErr) {
	zap.L().Info("Starting user creation service")

	emailExists := srv.db.GetUser(request.Email)

	if emailExists.Email != "" {
		zap.L().Warn("Email already associated with an existing account", zap.String("email", request.Email))
		return nil, exceptions.NewBadRequestError("Email is already associated with an existing account")
	}

	hashedPassword, err := srv.crypto.HashPassword(request.Password)
	if err != nil {
		zap.L().Error("Error when hashing password", zap.Error(err))
		return nil, exceptions.NewInternalServerError("Internal server error")
	}

	user := buildUserEntity(request, hashedPassword)

	dbErr := srv.db.CreateUser(*user)
	if dbErr != nil {
		zap.L().Error("Error creating user in database", zap.Error(dbErr))
		return nil, exceptions.NewInternalServerError("Internal server error")
	}

	zap.L().Info("User created successfully", zap.String("email", user.Email))
	return user, nil
}
