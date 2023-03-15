package domains

import (
	entities "mygram/domains/user/entity"
	"mygram/domains/user/model"
)

// Repository contract

type UserRepository interface {
	GetUserByEmail() 
	Insert(user entities.User) error
}

// Service contract
type UserUsecase interface {
	RegisterUser(model.RegisterUserRequest)
	FetchUserLogin()
}
