//go:build wireinject
// +build wireinject

package main

import (
	"mygram/applications/usecase"
	"mygram/commons/exceptions"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"mygram/infrastructures/validation"
	handler "mygram/interfaces/http/api"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)


var userSet = wire.NewSet(repository.NewUserRepositoryPostgres,usecase.NewUserUseCase,handler.NewUserHandler)
var photoSet = wire.NewSet(repository.NewPhotoRepository,usecase.NewPhotoUsecase,handler.NewPhotoHandler)
var commentSet = wire.NewSet(repository.NewCommentRepository,usecase.NewCommmentUsecase,handler.NewCommentHandler)
var socialmediaSet = wire.NewSet(repository.NewSocialmediaRepository,usecase.NewSocialmediaUsecase,handler.NewSocialmediaHandler)
var roleSet = wire.NewSet(repository.NewRoleRepository,usecase.NewRoleUsecase,handler.NewRoleHandler)
var permissionSet = wire.NewSet(repository.NewPermissionRepository,usecase.NewPermissionUsecase,handler.NewPermissionHandler)

func InitializeApp(filenames ...string) *fiber.App {
	wire.Build(
		NewServer,
		config.New,
		dbConfig.NewPostgresDB,
		validation.NewValidation,
		userSet,
		photoSet,
		commentSet,
		socialmediaSet,
		permissionSet,
		roleSet,
	)
	return nil
}

func NewServer(
	userHandler handler.UserHandler,
	photoHandler handler.PhotoHandler,
	commentHandler handler.CommentHandler,
	socialmediaHandler handler.SocialMediaHandler,
	permissionHandler handler.PermissionHandler,
	roleHandler handler.RoleHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	userHandler.Route(app)
	photoHandler.Route(app)
	commentHandler.Route(app)
	socialmediaHandler.Route(app)
	permissionHandler.Route(app)
	roleHandler.Route(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	
	return app
}