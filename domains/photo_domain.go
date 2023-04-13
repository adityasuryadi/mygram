package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

// repository contract
type PhotoRepository interface {
	InsertPhoto(entities.Photo) (*entities.Photo,error)
	GetAll() ([]*entities.Photo,error)
	FindById(id string)(*entities.Photo,error)
	UpdatePhoto(id string,photo entities.Photo) error
	DestroyPhoto(id string) error
}

// service contract
type PhotoUsecase interface {
	CreatePhoto(model.CreatePhotoRequest) (interface{},string)
	FindAll() ([]model.PhotoResponse,string)
	GetPhotoById(ctx *fiber.Ctx,id string) (*model.PhotoResponse,string)
	EditPhoto(ctx *fiber.Ctx) (interface{},string)
	DeletePhoto(ctx *fiber.Ctx) (string)
}
