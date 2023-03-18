package model

type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,unique=user"`
	Email    string `json:"email" validate:"required,email,unique=user"`
	Password string `json:"password" validate:"required"`
	Age      int    `json:"age"`
}