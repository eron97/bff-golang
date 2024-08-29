package persistence_test

import (
	"testing"
	"time"

	entity "github.com/eron97/bff-golang.git/src/model/entitys"
	"github.com/eron97/bff-golang.git/src/model/persistence"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestGetUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entity.CreateUser{})

	repo := persistence.NewDBConnection(db)

	user := entity.CreateUser{
		ID:         entity.NewID(),
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		CepBR:      "12345678",
		Country:    "Brazil",
		City:       "São Paulo",
		Address:    "Rua das Flores, 123",
		Password:   "password",
		CreateAt:   time.Now(),
	}

	db.Create(&user)

	retrievedUser := repo.GetUser(user.Email)

	assert.NotNil(t, retrievedUser)
	assert.Equal(t, user.First_Name, retrievedUser.First_Name)
	assert.Equal(t, user.Last_Name, retrievedUser.Last_Name)
	assert.Equal(t, user.Email, retrievedUser.Email)
}

func TestCreateUser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entity.CreateUser{})

	repo := persistence.NewDBConnection(db)

	user := entity.CreateUser{
		ID:         entity.NewID(),
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		CepBR:      "12345678",
		Country:    "Brazil",
		City:       "São Paulo",
		Address:    "Rua das Flores, 123",
		Password:   "password",
		CreateAt:   time.Now(),
	}

	err = repo.CreateUser(user)
	assert.NoError(t, err)

	var createdUser entity.CreateUser

	result := db.First(&createdUser, "email = ?", user.Email)
	assert.NoError(t, result.Error)
	assert.Equal(t, user.First_Name, createdUser.First_Name)
	assert.Equal(t, user.Last_Name, createdUser.Last_Name)
	assert.Equal(t, user.Email, createdUser.Email)
}

func TestVerifyExist(t *testing.T) {

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&entity.CreateUser{})

	repo := persistence.NewDBConnection(db)

	user := entity.CreateUser{
		ID:         entity.NewID(),
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "john.doe@example.com",
		CepBR:      "12345678",
		Country:    "Brazil",
		City:       "São Paulo",
		Address:    "Rua das Flores, 123",
		Password:   "password",
		CreateAt:   time.Now(),
	}

	db.Create(&user)

	exists, err := repo.VerifyExist(user.Email)
	assert.NoError(t, err)
	assert.True(t, exists)

	exists, err = repo.VerifyExist("nonexistent@example.com")
	assert.NoError(t, err)
	assert.False(t, exists)
}
