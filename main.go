package main

import (
	"fmt"
	photoUsecase "mygram/applications/usecase/photo"
	usecase "mygram/applications/usecase/user"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"mygram/infrastructures/validation"
	handler "mygram/interfaces/http/api/users"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
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
	photoUsecase := photoUsecase.NewPhotoUsecase(photoRepository)
	PhotoHandler := 
	
	
	app.Get("/", func(c *fiber.Ctx) error {
		var total int64
		err := db.Table("user").Where("email = ?","adit@mail.com").Count(&total).Error
		// err := db.Raw("SELECT count(id) from user where email = adit@mail.com").Scan(&total)
		fmt.Println(total)
		if err != nil {
			fmt.Println(err)
		}
		return c.SendString("Hello, Aditya s!")
	})

	// Start App
	app.Listen(":5000")
	// if err != nil {
	// 	panic(err)
	// }
}