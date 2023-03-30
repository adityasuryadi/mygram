package main

import (
	usecase "mygram/applications/usecase"
	"mygram/commons/exceptions"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"mygram/infrastructures/validation"
	handler "mygram/interfaces/http/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	configApp := config.New()
	db:=dbConfig.NewPostgresDB(configApp)
	validate:=validation.NewValidation(db)
	
	// user
	userRepository:=repository.NewUserRepositoryPostgres(db)
	userUsecase:=usecase.NewUserUseCase(userRepository)
	userHandler:=handler.NewUserHandler(userUsecase,*validate)
	userHandler.Route(app)

	// photo
	photoRepository := repository.NewPhotoRepository(db)
	photoUsecase := usecase.NewPhotoUsecase(photoRepository,userRepository)
	photoHandler := handler.NewPhotoHandler(photoUsecase,*validate)
	photoHandler.Route(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	
	// Start App
	app.Listen(":5000")
}