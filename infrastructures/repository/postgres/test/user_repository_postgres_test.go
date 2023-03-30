package postgresql_repository_test

import (
	"fmt"
	entities "mygram/domains/entity"
	repository "mygram/infrastructures/repository/postgres"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
// var categoryService = CategoryService{Repository: categoryRepository}

// var userRepository = repository.UserRepositoryImpl

// func TestCategoryService_GetNotFound(t *testing.T) {
// 	categoryRepository.Mock.On("FindById", "1").Return(nil)
// 	category, err := categoryService.Get("1")
// 	assert.Nil(t, category)
// 	assert.NotNil(t, err)
// }

// func TestCategoryService_GetSuccess(t *testing.T) {
// 	category := entity.Category{
// 		Id:   "1",
// 		Name: "Laptop",
// 	}
// 	categoryRepository.Mock.On("FindById", "2").Return(category)
// 	result, err := categoryService.Get("2")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, category.Id, result.Id)
// 	assert.Equal(t, category.Name, result.Name)
// }

func TestUserRepository_SuccessGetUserByEmail(t *testing.T){
	// var userRepositoryMock = &mock_repository.UserRepositoryMock{
	// 	Mock: mock.Mock{},
	// }
	db, mock, err := sqlmock.New()
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

	user:=entities.User{
		Id:        uuid.New(),
		UserName:  "Adit",
		Email:     "adit@mail.com",
		Password:  "12345",
		Age:       21,
	}

	userRepository := repository.NewUserRepositoryPostgres(gormDB)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "email" = $1 ORDER BY "user"."id" LIMIT 1`)).
	WithArgs("adit@mail.com").
	WillReturnRows(sqlmock.
		NewRows([]string{"id","username","email","age","password"}).
		AddRow(user.Id.String(),user.UserName,user.Email,user.Age,user.Password),
	)
	// userRepositoryMock.Mock.On("GetUserByEmail","1").Return(user)
	result,err:=userRepository.GetUserByEmail("adit@mail.com")
	assert.NotNil(t,result)
	assert.Nil(t,err)
}

func TestUserRepository_NotFoundGetUserByEmail(t *testing.T){
	// var userRepositoryMock = &mock_repository.UserRepositoryMock{
	// 	Mock: mock.Mock{},
	// }
	db, mock, err := sqlmock.New()
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

	userRepository := repository.NewUserRepositoryPostgres(gormDB)
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "user" WHERE "email" = $1 ORDER BY "user"."id" LIMIT 1`)).
	WithArgs("adit@mail.com").
	WillReturnRows(sqlmock.
		NewRows([]string{"id"}),
	)
	result,err:=userRepository.GetUserByEmail("adit@mail.com")
	assert.Nil(t,result)
	assert.NotNil(t,err)
}