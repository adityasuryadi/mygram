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
}

func (handler SocialMediaHandler) PostSocialmedia(ctx *fiber.Ctx) error {
	var request model.CreateSocialmediaRequest
	
	
	// user := ctx.Locals("user").(*jwt.Token)
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