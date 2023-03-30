package usecase

import (
	"fmt"
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

func NewPhotoUsecase(repository domains.PhotoRepository,userRepository domains.UserRepository) domains.PhotoUsecase {
	return &PhotoUSecaseImpl{
		repository: repository,
		UserRepository: userRepository,
	}
}

type PhotoUSecaseImpl struct {
	repository domains.PhotoRepository
	UserRepository domains.UserRepository
}

func (usecase PhotoUSecaseImpl) CreatePhoto(request model.CreatePhotoRequest) (*model.CreatePhotoResponse,string) {
	var response model.CreatePhotoResponse
	
	errorCode := make(chan string,1)
	
	userResult,err := usecase.UserRepository.GetUserByEmail(request.Email)
	fmt.Println(userResult)
	if err != nil {
		errorCode <- "500"
		response = model.CreatePhotoResponse{}
	}
	photo:=entities.Photo{
		Title:     request.Title,
		Caption:   request.Caption, 
		PhotoUrl:  request.PhotoUrl,
		UserId: userResult.Id,
	}

	result,err := usecase.repository.InsertPhoto(photo)

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
			UpdatedAt: result.UpdatedAt,
		}
		errorCode <- "200"
	}

	return &response,<-errorCode
}
