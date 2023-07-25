package test_usecase

import (
	"fmt"
	"testing"

	"mygram/applications/usecase"
	"mygram/domains/model"
	mock_repository "mygram/infrastructures/repository/postgres/mock"
	"mygram/infrastructures/validation"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetUserBYEmail(t *testing.T) {
	userRepository := &mock_repository.UserRepositoryMock{
		Mock: mock.Mock{},
	}

	db, _, err := sqlmock.New()
	if err != nil {
		// t.Fatal("error creating mock database: %v",err)
		fmt.Println(err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("error opening GORM database: %v", err)
	}

	validation := validation.NewValidation(gormDB)

	userUsecase := usecase.NewUserUseCase(userRepository, validation)
	userRepository.Mock.On("GetUserByEmail", "adit@mail.com").Return(nil)
	userUsecase.RegisterUser(model.RegisterUserRequest{})
}
