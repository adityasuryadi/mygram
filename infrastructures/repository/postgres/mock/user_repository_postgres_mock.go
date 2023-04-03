package mock_repository

import (
	"errors"
	"fmt"
	entities "mygram/domains/entity"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	Mock mock.Mock 
}

func (repository *UserRepositoryMock) GetUserByEmail(email string) (*entities.User,error) {
	arguments:=repository.Mock.Called(email)
	if arguments.Get(0) == nil {
		return nil,errors.New("user notfound")
	}else{
		user := arguments.Get(0).(entities.User)
		return &user,nil  
	}
}

func (repository *UserRepositoryMock) Insert(user *entities.User) error {
	fmt.Println(user.Password)
	arguments := repository.Mock.Called(user)
	if arguments.Get(0) == nil {
		return errors.New("user notfound")
	}
	return nil
}