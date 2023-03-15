package handler

import (
	domains "mygram/domains/user"

	userModel "mygram/domains/user/model"

	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(usecase domains.UserUsecase) UserHandler{
	return UserHandler{UserUsecase: usecase}
}

type UserHandler struct {
	UserUsecase domains.UserUsecase
}

func (handler UserHandler) Route(app *fiber.App){
	app.Get("test",handler.GetUser)
	app.Post("register",handler.Register)
}

func (handler UserHandler) GetUser(ctx *fiber.Ctx) error {
	return ctx.SendString("ini test")
}

func (handler UserHandler) Register(ctx *fiber.Ctx)error{
	var request userModel.RegisterUserRequest
	ctx.BodyParser(&request)
	handler.UserUsecase.RegisterUser(request)
	return nil
}