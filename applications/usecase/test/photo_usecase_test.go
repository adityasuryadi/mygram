package test_usecase

import (
	"mygram/applications/usecase"
	"mygram/domains/model"
	mock_repository "mygram/infrastructures/repository/postgres/mock"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGetUserBYEmail(t *testing.T){
	var userRepository = &mock_repository.UserRepositoryMock{
		Mock: mock.Mock{},
	}
	userUsecase := usecase.NewUserUseCase(userRepository)
	userRepository.Mock.On("GetUserByEmail","adit@mail.com").Return(nil)
	userUsecase.RegisterUser(model.RegisterUserRequest{})

}