package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

type ProductRepository interface {
	Insert(product *entities.Product) error
	FindById(id string)(*entities.Product,error)
	Update(product *entities.Product) (*entities.Product,error)
	Delete(id string) error
}

type ProductUsecase interface {
	CreateProduct(ctx *fiber.Ctx) (string,interface{},*model.ProductResponse)
	FindProductById(ctx *fiber.Ctx,id string)(string,interface{},*model.ProductResponse)
	EditProduct(ctx *fiber.Ctx)(string,interface{},*model.ProductResponse)
	DeleteProduct(ctx *fiber.Ctx) (string,interface{},*model.ProductResponse)
}