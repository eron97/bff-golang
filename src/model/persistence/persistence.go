package persistence

import (
	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PersistenceInterface interface {
	CreateUser(user entity.CreateUser) error
	VerifyExist(email string) (bool, error)
	GetUser(email string) *entity.CreateUser
}

type DBConnection struct {
	db *gorm.DB
}

func NewDBConnection(db *gorm.DB) PersistenceInterface {
	return &DBConnection{db: db}
}

func (repo *DBConnection) CreateUser(user entity.CreateUser) error {
	zap.L().Info("Creating user in the database", zap.String("email", user.Email))
	err := repo.db.Create(&user).Error
	if err != nil {
		zap.L().Error("Error creating user in database", zap.Error(err))
	}
	return err
}

func (repo *DBConnection) VerifyExist(email string) (bool, error) {
	zap.L().Info("Checking user existence", zap.String("email", email))
	var count int64
	err := repo.db.Table("create_users").Where("email = ?", email).Count(&count).Error
	if err != nil {
		zap.L().Error("Error checking user existence", zap.Error(err))
	}
	return count > 0, err
}

func (repo *DBConnection) GetUser(email string) *entity.CreateUser {
	zap.L().Info("Getting user from database", zap.String("email", email))
	var user entity.CreateUser
	err := repo.db.Table("create_users").Where("email = ?", email).First(&user).Error
	if err != nil {
		zap.L().Error("User not found in database", zap.Error(err))
	}
	return &user
}
