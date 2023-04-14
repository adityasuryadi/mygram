package usecase

import (
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

func NewRoleUsecase(roleRepo domains.RoleRepository) domains.RoleUsecase{
	return &RoleUsecaseImpl{
		RoleRepo: roleRepo,
	}
}

type RoleUsecaseImpl struct {
	RoleRepo domains.RoleRepository
}

func (RoleUsecase *RoleUsecaseImpl) UpdateRole(ctx *fiber.Ctx)(string,interface{},*model.ResponseRole){
	responseCode := make(chan string,1)
	id:=ctx.Params("id")
	var request model.UpdateRequestRole
	ctx.BodyParser(&request)
	permissionInput := request.Permissions
	var permissions []entities.Permission


	for _, v := range permissionInput {
		permissions = append(permissions, entities.Permission{
			Id: int(v),
		})		
	}
	role := &entities.Role{
		Name: request.Name,
		Permissions: permissions,
	}
	responseCode <- "200"
	RoleUsecase.RoleRepo.Edit(id,role)
	return <-responseCode,nil,nil
}