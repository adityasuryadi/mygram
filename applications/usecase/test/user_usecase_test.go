package test_usecase

import (
	"database/sql/driver"
	"fmt"
	"mygram/applications/usecase"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	mock_repository "mygram/infrastructures/repository/postgres/mock"
	"mygram/infrastructures/validation"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestUserUsecase_RegisterUserSuccess(t *testing.T) {
	var userRepository = &mock_repository.UserRepositoryMock{
		Mock: mock.Mock{},
	}

	userRequest := model.RegisterUserRequest{
		Username: "aditya",
		Email:    "aditya@mail.com",
		Password: "12345",
		Age:      25,
	}

	// hashedPassword := security.GetHash([]byte(userRequest.Password))

	user:=&entities.User{
		UserName:  userRequest.Username,
		Email:     userRequest.Email,
		Password: userRequest.Password,
		Age:       userRequest.Age,
	}

	db, _, err := sqlmock.New()
	if err != nil {
		// t.Fatal("error creating mock database: %v",err)
		fmt.Println(err)
	}

	gormDB,err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}),&gorm.Config{})

	if err != nil {
        t.Fatalf("error opening GORM database: %v", err)
    }

	validation:= validation.NewValidation(gormDB)

	userRepository.Mock.On("Insert",user).Return(user)
	userUsecase := usecase.NewUserUseCase(userRepository,validation)
	errorCode := userUsecase.RegisterUser(userRequest)
	assert.Equal(t,errorCode,"200")

	userRepository.Mock.AssertCalled(t,"Insert",user)
}