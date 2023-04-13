package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/infrastructures/security"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewSocialmediaHandler(socialMediaUsecase domains.SocialmediaUsecase) SocialMediaHandler {
	return SocialMediaHandler{
		usecase: socialMediaUsecase,
	}
}

type SocialMediaHandler struct {
	usecase domains.SocialmediaUsecase
}

func (handler SocialMediaHandler) Route(app *fiber.App){
	socialmedia := app.Group("socialmedia",middleware.Verify())
	socialmedia.Post("/",handler.PostSocialmedia)
	socialmedia.Get("/",handler.GetAllSocialmedia)
	socialmedia.Get("/:id",handler.GetOneSocialmedia)
	socialmedia.Put("/:id",handler.EditSocialmedia)
	socialmedia.Delete("/:id",handler.DeleteSocialmedia)
}

func (handler SocialMediaHandler) PostSocialmedia(ctx *fiber.Ctx) error {
	var request model.CreateSocialmediaRequest
	
	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)
	request.Email = email
	ctx.BodyParser(&request)

	responseCode,validation,data := handler.usecase.CreateSocialmedia(request)
	
	if responseCode == "400" {
		model.GetResponse(ctx,responseCode,"",validation)
		return nil
	}

	model.GetResponse(ctx,responseCode,"",data)
	return nil
}

func (handler SocialMediaHandler) GetAllSocialmedia(ctx *fiber.Ctx) error {
	responseCode,_,data := handler.usecase.ListSocialmedia()
	model.GetResponse(ctx,responseCode,"",data)
	return nil
}


func (handler SocialMediaHandler) GetOneSocialmedia(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	responseCode,_,data := handler.usecase.FindSocialmediaById(id)
	model.GetResponse(ctx,responseCode,"",data)
	return nil
}

func (handler SocialMediaHandler) EditSocialmedia(ctx *fiber.Ctx) error{
	var request model.CreateSocialmediaRequest
	ctx.BodyParser(&request)

	responseCode,validation,response := handler.usecase.EditSocialmedia(ctx)
	if responseCode == "400" {
		model.GetResponse(ctx,responseCode,"",validation)
		return nil
	}

	model.GetResponse(ctx,responseCode,"",response)
	return nil
}

func (handler SocialMediaHandler) DeleteSocialmedia(ctx *fiber.Ctx) error {
	responseCode,_,response := handler.usecase.DeleteSocialmedia(ctx)
	model.GetResponse(ctx,responseCode,"",response)
	return nil
}