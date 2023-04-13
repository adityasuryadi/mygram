package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

type SocialmediaRepository interface {
	Insert(entities.Socialmedia) error
	GetAll() ([]*entities.Socialmedia,error)
	FindById(id string)(*entities.Socialmedia,error)
	Update(socialmedia *entities.Socialmedia)(*entities.Socialmedia,error)
	Delete(id string) error
}

type SocialmediaUsecase interface {
	CreateSocialmedia(model.CreateSocialmediaRequest) (string,interface{},*model.SocialmediaResponse)
	ListSocialmedia()(string,interface{},[]*model.SocialmediaResponse)
	FindSocialmediaById(id string)(string,interface{},*model.SocialmediaResponse)
	EditSocialmedia(ctx *fiber.Ctx)(string,interface{},*model.SocialmediaResponse)
	DeleteSocialmedia(ctx *fiber.Ctx) (string,interface{},*model.SocialmediaResponse)
}
