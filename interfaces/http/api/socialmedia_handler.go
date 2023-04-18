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


// Create a new socialmedia
// @Summary Register a new socialmedia
// @Description Register socialmedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Param socialmedia body model.CreateSocialmediaRequest true "Create Social Media"
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /socialmedia/ [post]
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

// GetAllSocialmedia list socialmedia
// @Summary list socialmedia
// @Description list socialmedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.SocialmediaResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /socialmedia/ [get]
func (handler SocialMediaHandler) GetAllSocialmedia(ctx *fiber.Ctx) error {
	responseCode,_,data := handler.usecase.ListSocialmedia()
	model.GetResponse(ctx,responseCode,"",data)
	return nil
}

// FindSocialmedia find socialmedia by id
// @Summary find socialmedia
// @Description find socialmedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=model.SocialmediaResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /socialmedia/{id} [GET]
func (handler SocialMediaHandler) GetOneSocialmedia(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	responseCode,_,data := handler.usecase.FindSocialmediaById(id)
	model.GetResponse(ctx,responseCode,"",data)
	return nil
}

// Edit Socialmedia
// @Summary edit socialmedia
// @Description Edit socialmedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Param user body model.CreateSocialmediaRequest true "Update socialmedia"
// @Success 200 {object} model.WebResponse{data=model.SocialmediaResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /socialmedia/{id} [PUT]
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

// Delete socialmedia
// @Summary Delete Socialmedia
// @Description Delete socialmedia
// @Tags socialmedia
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /socialmedia/{id} [DELETE]
func (handler SocialMediaHandler) DeleteSocialmedia(ctx *fiber.Ctx) error {
	responseCode,_,response := handler.usecase.DeleteSocialmedia(ctx)
	model.GetResponse(ctx,responseCode,"",response)
	return nil
}