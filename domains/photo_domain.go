package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

// repository contract
type PhotoRepository interface {
	InsertPhoto(entities.Photo) (*entities.Photo,error)
}

// service contract
type PhotoUsecase interface {
	CreatePhoto(model.CreatePhotoRequest) (*model.CreatePhotoResponse,string)
}