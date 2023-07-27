package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/infrastructures/security"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewProductHandler(usecase domains.ProductUsecase) ProductHandler {
	return ProductHandler{
		usecase: usecase,
	}
}

func (handler ProductHandler) Route(app *fiber.App){
	product := app.Group("product",middleware.Verify())
	product.Post("/",handler.PostProduct)
	product.Get("/:id",handler.GetOneProduct)
	product.Put("/:id",handler.EditProduct)
	product.Delete("/:id",handler.DeleteProduct)
}

type ProductHandler struct {
	usecase domains.ProductUsecase
}

// Create a new Product
// @Summary Register a new Product
// @Description Register Product
// @Tags Product
// @Accept json
// @Produce json
// @Param Product body model.ProductCreateRequest true "Create Product"
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /product/ [post]
// @Security BearerAuth
func (handler ProductHandler) PostProduct(ctx *fiber.Ctx) error {
	responseCode,_,response := handler.usecase.CreateProduct(ctx)
	model.GetResponse(ctx,responseCode,"",response)
	return nil
}


// FindProduct find Product by id
// @Summary find Product
// @Description find Product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=model.ProductResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /product/{id} [GET]
// @Security BearerAuth
func (handler ProductHandler) GetOneProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)
	ctx.Locals("email",email)
	responseCode,_,data := handler.usecase.FindProductById(ctx,id)
	model.GetResponse(ctx,responseCode,"",data)
	return nil
}

// Edit Product
// @Summary edit Product
// @Description Edit Product
// @Tags Product
// @Accept json
// @Produce json
// @Param user body model.ProductCreateRequest true "Update Product"
// @Success 200 {object} model.WebResponse{data=model.ProductResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /product/{id} [PUT]
// @Security BearerAuth
func (handler ProductHandler) EditProduct(ctx *fiber.Ctx) error{
	var request model.ProductCreateRequest
	ctx.BodyParser(&request)

	responseCode,validation,response := handler.usecase.EditProduct(ctx)
	if responseCode == "400" {
		model.GetResponse(ctx,responseCode,"",validation)
		return nil
	}

	model.GetResponse(ctx,responseCode,"",response)
	return nil
}

// Delete Product
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /product/{id} [DELETE]
// @Security BearerAuth
func (handler ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	responseCode,_,response := handler.usecase.DeleteProduct(ctx)
	model.GetResponse(ctx,responseCode,"",response)
	return nil
}