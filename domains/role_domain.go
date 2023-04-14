package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

// repository contract
type RoleRepository interface {
	Insert(role *entities.Role) error
	Edit(id string,role *entities.Role) error
}

// usecase contract

type RoleUsecase interface {
	UpdateRole(ctx *fiber.Ctx) (string,interface{},*model.ResponseRole)
}