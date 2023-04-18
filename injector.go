//go:build wireinject
// +build wireinject

package main

import (
	"mygram/applications/usecase"
	config "mygram/infrastructures"
	dbConfig "mygram/infrastructures/database"
	repository "mygram/infrastructures/repository/postgres"
	"mygram/infrastructures/validation"
	handler "mygram/interfaces/http/api"

	"github.com/google/wire"
)

var userSet = wire.NewSet(repository.NewUserRepositoryPostgres,usecase.NewUserUseCase,handler.NewUserHandler)
var photoSet = wire.NewSet(repository.NewPhotoRepository,repository.NewUserRepositoryPostgres,usecase.NewPhotoUsecase,handler.NewPhotoHandler)
var commentSet = wire.NewSet(repository.NewCommentRepository,repository.NewUserRepositoryPostgres,usecase.NewCommmentUsecase,handler.NewCommentHandler)
var socialmediaSet = wire.NewSet(repository.NewSocialmediaRepository,repository.NewUserRepositoryPostgres,usecase.NewSocialmediaUsecase,handler.NewSocialmediaHandler)
var roleSet = wire.NewSet(repository.NewRoleRepository,usecase.NewRoleUsecase,handler.NewRoleHandler)
var permissionSet = wire.NewSet(repository.NewPermissionRepository,usecase.NewPermissionUsecase,handler.NewPermissionHandler)

func InitializedUserHandler(filenames ...string) handler.UserHandler{
	wire.Build(
		config.New,
		dbConfig.NewPostgresDB,
		validation.NewValidation,
		userSet,
	)
	return handler.UserHandler{}
}

func InitializedPhotoHandler(filenames ...string) handler.PhotoHandler{
	wire.Build(
		config.New,
		dbConfig.NewPostgresDB,
		validation.NewValidation,
		photoSet,
	)
	return handler.PhotoHandler{}
}

func InitializedCommentHandler(filenames ...string) handler.CommentHandler{
	wire.Build(
		config.New,
		dbConfig.NewPostgresDB,
		validation.NewValidation,
		commentSet,
	)
	return handler.CommentHandler{}
}

func InitializedSocialmediaHandler(filenames ...string) handler.SocialMediaHandler{
	wire.Build(
		config.New,
		dbConfig.NewPostgresDB,
		validation.NewValidation,
		socialmediaSet,
	)
	return handler.SocialMediaHandler{}
}

func InitializedRoleHandler(filenames ...string) handler.RoleHandler{
	wire.Build(
		config.New,
		dbConfig.NewPostgresDB,
		roleSet,
	)
	return handler.RoleHandler{}
}

func InitializedPermissionHandler(filenames ...string) handler.PermissionHandler{
	wire.Build(
		config.New,
		dbConfig.NewPostgresDB,
		permissionSet,
	)
	return handler.PermissionHandler{}
}