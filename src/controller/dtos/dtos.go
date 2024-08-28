package dtos

type CreateUser struct {
	First_Name string `json:"first_name" validate:"required,min=4,max=100" example:"John"`
	Last_Name  string `json:"last_name" validate:"required,min=4,max=100" example:"Doe"`
	Email      string `json:"email" validate:"required,email" example:"test@test.com"`
	CepBR      string `json:"cep" validate:"required,len=8" example:"12345678"`
	Country    string `json:"country" validate:"required" example:"Brazil"`
	City       string `json:"city" validate:"required" example:"São Paulo"`
	Address    string `json:"address" validate:"required,min=5,max=200" example:"Rua das Flores, 123"`
	Password   string `json:"password" validate:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email" example:"test@test.com"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
}

type NewUser struct {
	First_Name string
	Email      string
}
