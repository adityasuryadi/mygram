package main

import (
	usecase "mygram/applications/usecase/user"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	handler "mygram/interfaces/http/api/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	configApp := config.New()
	db:=dbConfig.NewPostgresDB(configApp)
	userRepository:=repository.NewUserRepositoryPostgres(db)
	userUsecase:=usecase.NewUserUseCase(userRepository)
	userHandler:=handler.NewUserHandler(userUsecase)
	userHandler.Route(app)
	
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Aditya s!")
	})

	// Start App
	err := app.Listen(":5000")
	if err != nil {
		panic(err)
	}
}