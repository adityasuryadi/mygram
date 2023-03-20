package usecase

import (
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

func NewPhotoUsecase(repository domains.PhotoRepository) domains.PhotoUsecase {
	return &PhotoUSecaseImpl{
		repository: repository,
	}
}

type PhotoUSecaseImpl struct {
	repository domains.PhotoRepository
}

func (usecase PhotoUSecaseImpl) CreatePhoto(request model.CreatePhotoRequest) (*model.CreatePhotoResponse,string) {
	var response model.CreatePhotoResponse
	photo:=entities.Photo{
		Title:     request.Title,
		Caption:   request.Caption,
		PhotoUrl:  request.PhotoUrl,
		UserId:    [16]byte{},
	}
	result,err := usecase.repository.InsertPhoto(photo)
	errorCode := make(chan string,1)
	if err != nil {
		errorCode <- "500"
		response = model.CreatePhotoResponse{}
	}

	if err == nil {
		response = model.CreatePhotoResponse{
			Id:        result.Id,
			PhotoUrl:  result.PhotoUrl,
			Caption:   result.Caption,
			Title: result.Title,
			CreatedAt: result.CreatedAt,
			UodatedAt: result.UpdatedAt,
		}
		errorCode <- "200"
	}

	return &response,<-errorCode
}
