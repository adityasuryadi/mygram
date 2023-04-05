package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"
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
	GetPhotoById(id string) (*model.PhotoResponse,string)
	EditPhoto(id string,request model.UpdatePhotoRequest) (interface{},string)
	DeletePhoto(id string) (string)
}
