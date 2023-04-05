package usecase

import (
	"fmt"
	domains "mygram/domains"
	userEntities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/security"
	"mygram/infrastructures/validation"
	"reflect"
)

func NewUserUseCase(repository domains.UserRepository,validate validation.Validation) domains.UserUsecase {
	return &UserUseCaseImpl{
		repository: repository,
		Validate: validate,
	}
}

type UserUseCaseImpl struct {
	repository domains.UserRepository
	Validate validation.Validation
}

// RegisterUser implements domains.UserUsecase
func (usecase *UserUseCaseImpl) RegisterUser(request model.RegisterUserRequest) (errorCode string) {
	user := &userEntities.User{
		UserName:  request.Username,
		Email:     request.Email,
		Password:  security.GetHash([]byte(request.Password)),
		// Password: request.Password,
		Age:       request.Age,
	}
	err:=usecase.repository.Insert(user)
	if err != nil {
		return "500"
	}
	return "200"
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
