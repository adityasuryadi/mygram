package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/infrastructures/helper"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewPhotoHandler(usecase domains.PhotoUsecase) PhotoHandler {
	return PhotoHandler{
		usecase: usecase,
	}
}

type PhotoHandler struct {
	usecase domains.PhotoUsecase
}

func (handler PhotoHandler) Route(app *fiber.App){	
	// auth
	photo:=app.Group("photo",middleware.Verify())
	photo.Post("/",handler.PostPhoto)	
	photo.Get("/",handler.ListPhoto)
	photo.Get("/:id",handler.GetPhoto)
	photo.Put("/:id",handler.UpdatePhoto)
	photo.Delete("/:id",handler.DeletePhoto)
}

// PostPhoto godoc
// @Summary Create Photo
// @Description Create Photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photo body model.CreatePhotoRequest true "Create Photo"
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {string} model.WebResponse{code=400}
// @Failure 500 {string} model.WebResponse{code=500}
// @Router /photo [post]
// @Security BearerAuth
func (handler PhotoHandler) PostPhoto(ctx *fiber.Ctx) error {
	var request model.CreatePhotoRequest
	ctx.BodyParser(&request)
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	can,_:=helper.Can(email,"photo.create")
	
	if !can {
		model.ForbiddenResponse(ctx,"FORBIDDEN",nil)
		return nil
	}
	
	// request.Email = email

	result,errCode := handler.usecase.CreatePhoto(request)
	model.GetResponse(ctx,errCode,"",result)
	return nil
}


// Registeruser List Photo 
// @Summary List Photo
// @Description List Photo
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.PhotoResponse}
// @Failure 400 {string} model.WebResponse{code=400}
// @Failure 500 {string} model.WebResponse{code=500}
// @Router /photo [GET]
// @Security BearerAuth
func (handler PhotoHandler) ListPhoto(ctx *fiber.Ctx) error{
	res,errCode := handler.usecase.FindAll()
	model.GetResponse(ctx,errCode,"",res)
	return nil
}


// Registeruser List Photo 
// @Summary List Photo
// @Description List Photo
// @Tags photo
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=model.PhotoResponse}
// @Failure 404 {string} model.WebResponse{code=404}
// @Failure 500 {string} model.WebResponse{code=500}
// @Router /photo/id [GET]
// @Security BearerAuth
func (handler PhotoHandler) GetPhoto(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	
	res,responseCode := handler.usecase.GetPhotoById(ctx,id)

	model.GetResponse(ctx,responseCode,"",res)

	return nil
}

// Registeruser Edit Photo 
// @Summary Edit Photo
// @Description Edit Photo
// @Tags photo
// @Accept json
// @Produce json
// @Param photo body model.UpdatePhotoRequest true "Update photo"
// @Param id path string true "Photo ID"
// @Success 200 {object} model.WebResponse{data=model.PhotoResponse}
// @Failure 404 {string} model.WebResponse{code=404}
// @Failure 500 {string} model.WebResponse{code=500}
// @Router /photo/id [PUT]
// @Security BearerAuth
func (handler PhotoHandler) UpdatePhoto(ctx *fiber.Ctx) error {
	result,errCode := handler.usecase.EditPhoto(ctx)
	model.GetResponse(ctx,errCode,"",result)
	return nil
}

// DeletePhoto function removes a photo by ID
// @Summary Remove photo by ID
// @Description Remove photo by ID
// @Tags photo
// @Accept json
// @Produce json
// @Param id path string true "Photo ID"
// @Success 200 {object} model.WebResponse{}
// @Failure 404 {object} model.WebResponse{}
// @Failure 503 {object} model.WebResponse{}
// @Router /photo/{id} [delete]
// @Security BearerAuth
func (handler PhotoHandler) DeletePhoto(ctx *fiber.Ctx) error {
	errCode := handler.usecase.DeletePhoto(ctx)
	model.GetResponse(ctx,errCode,"",nil)
	return nil
}