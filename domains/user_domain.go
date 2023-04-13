package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

// Repository contract

type UserRepository interface {
	GetUserByEmail(email string) (*entities.User,error) 
	Insert(user *entities.User) error
	AssignRole(userId string,roles []int) error
}

// Service contract
type UserUsecase interface {
	RegisterUser(request model.RegisterUserRequest) (string,interface{})
	FetchUserLogin(model.LoginUserRequest) (string,string)
}

