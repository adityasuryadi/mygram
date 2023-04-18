package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewRoleHandler(roleUsecase domains.RoleUsecase) RoleHandler{
	return RoleHandler{
		RoleUsecase: roleUsecase,
	}
}

type RoleHandler struct{
	RoleUsecase domains.RoleUsecase
}

func (handler RoleHandler) Route(app *fiber.App){
	role:=app.Group("role",middleware.Verify())
	role.Put("/:id",handler.PutRole)
}


func (handler RoleHandler) PutRole(ctx *fiber.Ctx)error{
	responseCode,_,result := handler.RoleUsecase.UpdateRole(ctx)
	model.GetResponse(ctx,responseCode,"",result)
	return nil
}
