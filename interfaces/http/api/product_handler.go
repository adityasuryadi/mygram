package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewProductHandler(usecase domains.ProductUsecase) ProductHandler {
	return ProductHandler{
		usecase: usecase,
	}
}

func (handler ProductHandler) Route(app *fiber.App){
	product := app.Group("product",middleware.Verify())
	product.Post("product",handler.PostProduct)
}

type ProductHandler struct {
	usecase domains.ProductUsecase
}

func (handler ProductHandler) PostProduct(ctx *fiber.Ctx) error {
	responseCode,_,response := handler.usecase.CreateProduct(ctx)
	model.GetResponse(ctx,responseCode,"",response)
	return nil
}