package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

type ProductRepository interface {
	Insert(product *entities.Product) error
}

type ProductUsecase interface {
	CreateProduct(ctx *fiber.Ctx) (string,interface{},*model.ProductResponse)
}