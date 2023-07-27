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

var (
	userSet        = wire.NewSet(repository.NewUserRepositoryPostgres, repository.NewUserTokenRepository, usecase.NewUserUseCase, handler.NewUserHandler)
	photoSet       = wire.NewSet(repository.NewPhotoRepository, usecase.NewPhotoUsecase, handler.NewPhotoHandler)
	commentSet     = wire.NewSet(repository.NewCommentRepository, usecase.NewCommmentUsecase, handler.NewCommentHandler)
	socialmediaSet = wire.NewSet(repository.NewSocialmediaRepository, usecase.NewSocialmediaUsecase, handler.NewSocialmediaHandler)
	roleSet        = wire.NewSet(repository.NewRoleRepository, usecase.NewRoleUsecase, handler.NewRoleHandler)
	permissionSet  = wire.NewSet(repository.NewPermissionRepository, usecase.NewPermissionUsecase, handler.NewPermissionHandler)
	productSet     = wire.NewSet(repository.NewProductRepository, usecase.NewProductUsecase, handler.NewProductHandler)
	fileSet        = wire.NewSet(usecase.NewFileUsecase, handler.NewFileHandler)
)

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
		productSet,
		fileSet,
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
	productHandler handler.ProductHandler,
	fileHandler handler.FileHandler,
) *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: exceptions.ErrorHandler})
	userHandler.Route(app)
	photoHandler.Route(app)
	commentHandler.Route(app)
	socialmediaHandler.Route(app)
	permissionHandler.Route(app)
	roleHandler.Route(app)
	productHandler.Route(app)
	fileHandler.Route(app)
	return app
}
