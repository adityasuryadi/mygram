package test_usecase

import (
	"mygram/applications/usecase"
	entities "mygram/domains/entity"
	mock_repository "mygram/infrastructures/repository/postgres/mock"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valyala/fasthttp"
)

var productRepo = &mock_repository.ProductRepositoryMock{Mock: mock.Mock{}}
var userRepo = &mock_repository.UserRepositoryMock{Mock: mock.Mock{}}
var productUsecase = usecase.NewProductUsecase(productRepo,userRepo)

func TestProductUsecase_FindProductByIdNotFound(t *testing.T) {
	user:=entities.User{
		Id: uuid.MustParse("33a61a3d-88e8-484d-8061-3db0bff92e3a"),
		Email: "adit@mail.com",
		Roles: []entities.Role{
			{Name: "user"},
		},
	}
	productRepo.Mock.On("FindById","1").Return(nil,nil)
	userRepo.Mock.On("GetUserByEmail","adit@mail.com").Return(user)
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	c.Context().SetUserValue("email","adit@mail.com")
	responseCode,_,result := productUsecase.FindProductById(c,"1")
	assert.Nil(t,result)
	assert.Equal(t,responseCode,"404")
}

func TestProductUsecase_FindProductByIdFound(t *testing.T) {
	user:=entities.User{
		Id: uuid.MustParse("33a61a3d-88e8-484d-8061-3db0bff92e3a"),
		Email: "adit@mail.com",
		Roles: []entities.Role{
			{Name: "user"},
		},
	}

	product:=entities.Product{
		Id:uuid.MustParse("69d028cb-bc49-4097-8554-3a1f742360c3"),
		Name: "Susu Kotak",
		Stock: 100,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserId: uuid.MustParse("33a61a3d-88e8-484d-8061-3db0bff92e3a"),

	}
	productRepo.Mock.On("FindById","69d028cb-bc49-4097-8554-3a1f742360c3").Return(product,nil)
	userRepo.Mock.On("GetUserByEmail","adit@mail.com").Return(user)
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	c.Context().SetUserValue("email","adit@mail.com")
	responseCode,_,result := productUsecase.FindProductById(c,"69d028cb-bc49-4097-8554-3a1f742360c3")
	assert.NotNil(t,result)
	assert.Equal(t,responseCode,"200")
}