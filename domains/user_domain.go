package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

// Repository contract

type UserRepository interface {
	GetUserByEmail(email string) (*entities.User, error)
	Insert(user *entities.User) error
	AssignRole(userId string, roles []entities.Role) error
}

// Service contract
type UserUsecase interface {
	RegisterUser(request model.RegisterUserRequest) (string, interface{})
	FetchUserLogin(model.LoginUserRequest) (string, string)
	UpdateUserRole(ctx *fiber.Ctx) (string, interface{}, interface{})
	Logout(ctx *fiber.Ctx) error
}
