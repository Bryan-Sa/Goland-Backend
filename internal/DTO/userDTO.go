package DTO

type UserLoginRequestBodyDTO struct {
	email    string `json:"email"  validate:"required,email"`
	password int    `json:"password"  validate:"required"`
}
