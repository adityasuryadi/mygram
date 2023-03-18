package handler

import (
	domains "mygram/domains/user"
	"mygram/infrastructures/validation"

	"mygram/domains/user/model"
	userModel "mygram/domains/user/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(usecase domains.UserUsecase,validate validator.Validate) UserHandler{
	return UserHandler{
		UserUsecase: usecase,
		validate:    validate,
	}
}

type UserHandler struct {
	UserUsecase domains.UserUsecase
	validate validator.Validate
}

func (handler UserHandler) Route(app *fiber.App){
	app.Get("test",handler.GetUser)
	app.Post("register",handler.Register)
}

func (handler UserHandler) GetUser(ctx *fiber.Ctx) error {
	return ctx.SendString("ini test")
}

func (handler UserHandler) Register(ctx *fiber.Ctx)error{
	// var validator *validator.Validate
	var request userModel.RegisterUserRequest
	validate := handler.validate
	ctx.BodyParser(&request)
	// // Register the unique validation function
	err := validate.Struct(&request)
	
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		out := make([]validation.ErrorMessage, len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = validation.ErrorMessage{fieldError.Field(), validation.GetErrorMsg(fieldError)}
		}
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(model.WebResponse{
				Code:   fiber.StatusBadRequest,
				Message: "FAILED TO CREATE DATA",
				Status: "BAD_REQUEST",
				Data:   out,
			})
	}

	handler.UserUsecase.RegisterUser(request)
	model.SuccessResponse(ctx,"SUCCESS CREATE DATA",nil)
	return nil
}