package handler

import (
	"fmt"
	domains "mygram/domains/user"
	"mygram/infrastructures/validation"
	"mygram/interfaces/http/api/middleware"

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
	app.Post("register",handler.Register)
	app.Post("login",handler.Login)
	app.Use(middleware.Verify())
	app.Get("test",handler.GetUser)
}

func (handler UserHandler) GetUser(ctx *fiber.Ctx) error {
	return ctx.SendString("ini test")
}

/*
REGISTER HANDLER
*/
func (handler UserHandler) Register(ctx *fiber.Ctx)error{
	var request userModel.RegisterUserRequest
	validate := handler.validate
	ctx.BodyParser(&request)
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

/* 
Login Handler
*/

func (handler UserHandler) Login(ctx *fiber.Ctx) error {
	var request userModel.LoginUserRequest
	validate := handler.validate
	ctx.BodyParser(&request)
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

	token,errorCode := handler.UserUsecase.FetchUserLogin(request)
	fmt.Println(errorCode)
	if errorCode == "404" {
		model.NotFoundResponse(ctx,"USER NOT FOUND",nil)
		return nil
	}
	
	if errorCode == "400" {
		model.BadRequestResponse(ctx,"WRONG EMAIL OR PASSWORD",nil)
		return nil
	}

	if errorCode == "500" {
		model.InternalServerErrorResponse(ctx,"SERVER FAILURE",nil)
		return nil
	}

	if errorCode == "200" {
		model.SuccessResponse(ctx,"SUCCESS LOGIN",model.LoginResponse{Token: token})
		return nil
	}

	return nil
}