package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

type SocialmediaRepository interface {
	Insert(entities.Socialmedia) error
}

type SocialmediaUsecase interface {
	CreateSocialmedia(model.CreateSocialmediaRequest) (string,interface{},*model.SocialmediaResponse)
}