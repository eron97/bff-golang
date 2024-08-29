package view

import (
	"github.com/eron97/bff-golang.git/src/controller/dtos"
	entity "github.com/eron97/bff-golang.git/src/model/entitys"
)

func ConvertDomainToResponse(resp *entity.CreateUser) dtos.NewUser {
	return dtos.NewUser{
		First_Name: resp.First_Name,
		Email:      resp.Email,
	}
}
