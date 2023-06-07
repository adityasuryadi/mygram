package routes

import (
	handler "mygram/interfaces/http/api"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewPhotoRoute(handler handler.PhotoHandler) *fiber.App{
	app := fiber.New()
	photo := app.Group("photo", middleware.Verify())
	photo.Post("/", handler.PostPhoto)
	photo.Get("/", handler.ListPhoto)
	photo.Get("/:id", handler.GetPhoto)
	photo.Put("/:id", handler.UpdatePhoto)
	photo.Delete("/:id", handler.DeletePhoto)

	return app
}