package usecase

import (
	"fmt"
	"mygram/commons/exceptions"
	domains "mygram/domains"
	userEntities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/security"
	"reflect"
)

func NewUserUseCase(repository domains.UserRepository) domains.UserUsecase {
	return &UserUseCaseImpl{
		repository: repository,
	}
}

type UserUseCaseImpl struct {
	repository domains.UserRepository
}

// RegisterUser implements domains.UserUsecase
func (usecase *UserUseCaseImpl) RegisterUser(request model.RegisterUserRequest) {
	user := userEntities.User{
		UserName:  request.Username,
		Email:     request.Email,
		Password:  security.GetHash([]byte(request.Password)),
		Age:       request.Age,
	}
	err:=usecase.repository.Insert(user)
	if err != nil {
		exceptions.PanicIfNeeded(err)
	}
}

func (usecase *UserUseCaseImpl) FetchUserLogin(request model.LoginUserRequest) (string,string) {
	email := request.Email
	result,err := usecase.repository.GetUserByEmail(email)
	
	errorCode := make(chan string, 1)
	var token string
	token = ""
	if reflect.ValueOf(result).IsZero() {
		fmt.Println(err)
		errorCode <- "404"
	}
	
	if !reflect.ValueOf(result).IsZero() {
		hashPassword := result.Password
		err = security.ComparePassword(hashPassword, request.Password)
		if err != nil {
			errorCode <- "400"
		}else{
			token,_ = security.ClaimToken(email)
			errorCode <- "200"
		}
	}	
	return token, <-errorCode
}
