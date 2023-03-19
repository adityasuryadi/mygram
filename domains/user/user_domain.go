package domains

import (
	entities "mygram/domains/user/entity"
	"mygram/domains/user/model"
)

// Repository contract

type UserRepository interface {
	GetUserByEmail(email string) (entities.User,error) 
	Insert(user entities.User) error
}

// Service contract
type UserUsecase interface {
	RegisterUser(request model.RegisterUserRequest)
	FetchUserLogin(model.LoginUserRequest) (string,string)
}

