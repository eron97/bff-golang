package entity

import (
	"time"

	"github.com/google/uuid"
)

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

type CreateUser struct {
	ID         ID
	First_Name string
	Last_Name  string
	Email      string
	CepBR      string
	Country    string
	City       string
	Address    string
	Password   string
	CreateAt   time.Time
}
