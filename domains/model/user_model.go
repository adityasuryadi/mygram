package model

// request
type RegisterUserRequest struct {
	Username string `json:"username" validate:"required,unique=user"`
	Email    string `json:"email" validate:"required,email,unique=user"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,gte=8"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

// response

type LoginResponse struct {
	Token string `json:"token"`
}