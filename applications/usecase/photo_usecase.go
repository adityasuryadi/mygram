package usecase

import (
	"fmt"
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	"time"
)

func NewPhotoUsecase(repository domains.PhotoRepository, userRepository domains.UserRepository) domains.PhotoUsecase {
	return &PhotoUSecaseImpl{
		repository:     repository,
		UserRepository: userRepository,
	}
}

type PhotoUSecaseImpl struct {
	repository     domains.PhotoRepository
	UserRepository domains.UserRepository
}

func (usecase PhotoUSecaseImpl) CreatePhoto(request model.CreatePhotoRequest) (*model.CreatePhotoResponse, string) {
	var response model.CreatePhotoResponse

	errorCode := make(chan string, 1)

	userResult, err := usecase.UserRepository.GetUserByEmail(request.Email)
	fmt.Println(userResult)
	if err != nil {
		errorCode <- "500"
		response = model.CreatePhotoResponse{}
	}
	photo := entities.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserId:   userResult.Id,
	}

	result, err := usecase.repository.InsertPhoto(photo)

	if err != nil {
		errorCode <- "500"
		response = model.CreatePhotoResponse{}
	}

	if err == nil {
		response = model.CreatePhotoResponse{
			Id:        result.Id,
			PhotoUrl:  result.PhotoUrl,
			Caption:   result.Caption,
			Title:     result.Title,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
		errorCode <- "200"
	}

	return &response, <-errorCode
}

// FindAll implements domains.PhotoUsecase
func (usecase PhotoUSecaseImpl) FindAll() ([]model.PhotoResponse, string) {
	var photos []model.PhotoResponse
	errorCode := make(chan string, 1)

	result, err := usecase.repository.GetAll()
	errorCode <- "200"
	if err != nil {
		errorCode <- "500"
	}

	for _, v := range result {
		photos = append(photos, model.PhotoResponse{
			Id:        v.Id,
			PhotoUrl:  v.PhotoUrl,
			Caption:   v.Caption,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return photos, <-errorCode
}

// GetPhotoById implements domains.PhotoUsecase
func (usecase *PhotoUSecaseImpl) GetPhotoById(id string) (*model.PhotoResponse, string) {
	errorCode := make(chan string, 1)
	result, err := usecase.repository.FindById(id)
	var response = &model.PhotoResponse{}

	if result == nil && err == nil {
		errorCode <- "404"
		response = nil
	}

	if err != nil && result == nil {
		errorCode <- "500"
		response = nil
	}

	if err == nil && result != nil {

		errorCode <- "200"
		response = &model.PhotoResponse{
			Id:        result.Id,
			PhotoUrl:  result.PhotoUrl,
			Caption:   result.Caption,
			Title:     result.Title,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
	}
	return response, <-errorCode
}

// EditPhoto implements domains.PhotoUsecase
func (usecase *PhotoUSecaseImpl) EditPhoto(id string, request model.UpdatePhotoRequest) (*model.UpdatePhotoResponse, string) {
	errCode := make(chan string, 1)
	var response = &model.UpdatePhotoResponse{}
	result, err := usecase.repository.FindById(id)
	
	if result == nil && err == nil {
		errCode <- "404"
		response = nil
		return response,<-errCode
	}

	if err != nil && result == nil {
		errCode <- "500"
		response = nil
		return response,<-errCode
	}

	result.PhotoUrl = request.PhotoUrl
	result.Caption = request.Caption
	result.Title = request.Title
	result.UpdatedAt = time.Now()

	err = usecase.repository.UpdatePhoto(id,*result)
	if err != nil {
		errCode <- "500"
		response = nil
		return response,<-errCode
	}
	errCode <- "200"
	return response,<-errCode
}
