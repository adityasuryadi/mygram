package main

import (
	"mygram/commons/exceptions"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	// docs are generated by Swag CLI, you have to import them.
	// replace with your own docs folder, usually "github.com/username/reponame/docs"
	_ "mygram/docs"
)

func NewApp() *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	return app
}

// @title TEST API
// @version 2.0
// @description This is a sample API Example
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host adisu.my.id
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// app := InitializeApp()
	// Start App
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	app.Listen(":5000")
}
