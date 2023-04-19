package usecase

import (
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

func NewProductUsecase(productRepository domains.ProductRepository,userRepo domains.UserRepository,) domains.ProductUsecase{
	return &ProductUsecaseImpl{
		productRepo: productRepository,
		userRepo: userRepo,
	}
}

type ProductUsecaseImpl struct {
	productRepo domains.ProductRepository
	userRepo domains.UserRepository	
}

func (productUsecase *ProductUsecaseImpl) CreateProduct(ctx *fiber.Ctx) (string,interface{},*model.ProductResponse) {
	var request model.ProductCreateRequest
	responseCode := make(chan string, 1)
	ctx.BodyParser(&request)
	product := entities.Product{
		Name: request.Name,
		Stock: request.Stock,
	}
	err := productUsecase.productRepo.Insert(&product)
	if err != nil {
		responseCode<-"500"
		return <-responseCode,nil,nil
	}
	responseCode <- "200"
	return <-responseCode,nil,nil
}