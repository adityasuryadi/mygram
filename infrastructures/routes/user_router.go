package routes

import (
	handler "mygram/interfaces/http/api"

	"github.com/gofiber/fiber/v2"
)

func NewUserRoute(app *fiber.App,handler handler.UserHandler) {
	app.Post("register",handler.Register)
	app.Post("login",handler.Login)
	app.Put("user/:id/assign",handler.PutUserRole)
}