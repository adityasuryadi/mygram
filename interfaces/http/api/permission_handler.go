package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewPermissionHandler(usecase domains.PermissionUsecase) *PermissionHandler {
	return  &PermissionHandler{
		usecase: usecase,
	}
}

type PermissionHandler struct {
	usecase domains.PermissionUsecase
}

func (handler PermissionHandler) Route(app *fiber.App){
	permission := app.Group("permission",middleware.Verify())
	permission.Post("/",handler.PostPermission)
}

func (handler PermissionHandler) PostPermission(ctx *fiber.Ctx) error {
	var request model.CreatePermissionRequest
	ctx.BodyParser(&request)
	responseCode,_,response := handler.usecase.CreatePermission(&request)
	model.GetResponse(ctx,responseCode,"",response)
	return nil
}