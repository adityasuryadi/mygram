package handler

import (
	"fmt"
	domains "mygram/domains"

	"mygram/domains/model"
	userModel "mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(usecase domains.UserUsecase) UserHandler{
	return UserHandler{
		UserUsecase: usecase,
	}
}

type UserHandler struct {
	UserUsecase domains.UserUsecase
}

func (handler UserHandler) Route(app *fiber.App){
	app.Post("register",handler.Register)
	app.Post("login",handler.Login)
}

func (handler UserHandler) GetUser(ctx *fiber.Ctx) error {
	return ctx.SendString("ini test")
}

/*
REGISTER HANDLER
*/

// Registeruser registers a new user data
// @Summary Register a new user
// @Description Register user
// @Tags users
// @Accept json
// @Produce json
// @Param user body userModel.RegisterUserRequest true "Register user"
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /register [post]
func (handler UserHandler) Register(ctx *fiber.Ctx)error{
	
	var request userModel.RegisterUserRequest
	ctx.BodyParser(&request)

	responseCode,data:=handler.UserUsecase.RegisterUser(request)
	model.GetResponse(ctx,responseCode,"",data)
	return nil
}

/* 
Login Handler
*/


// LoginuserLogin User
// @Summary Login user
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param user body userModel.LoginUserRequest true "Login user"
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /login [post]
func (handler UserHandler) Login(ctx *fiber.Ctx) error {
	var request userModel.LoginUserRequest
	ctx.BodyParser(&request)
	
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