package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/infrastructures/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewPhotoHandler(usecase domains.PhotoUsecase,validate validator.Validate) PhotoHandler {
	return PhotoHandler{
		usecase: usecase,
		validate: validate,
	}
}

type PhotoHandler struct {
	usecase domains.PhotoUsecase
	validate validator.Validate
}

func (handler PhotoHandler) Route(app *fiber.App){
	app.Post("photo",handler.PostPhoto)	
}

func (handler PhotoHandler) PostPhoto(ctx *fiber.Ctx) error {
	var request model.CreatePhotoRequest
	ctx.BodyParser(&request)
	err := handler.validate.Struct(&request)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		out := make([]validation.ErrorMessage,len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = validation.ErrorMessage{
				fieldError.Field(),
				validation.GetErrorMsg(fieldError),
			}
		}
		model.BadRequestResponse(ctx,"CLIENT SERVER ERROR",out)
		return nil
	}

	handler.usecase.CreatePhoto(request)
	model.SuccessResponse(ctx,"SUCCESS CREATE PHOTO",nil)
	return nil
}