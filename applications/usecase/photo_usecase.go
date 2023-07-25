package usecase

import (
	"fmt"
	"time"

	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/helper"
	"mygram/infrastructures/security"
	"mygram/infrastructures/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewPhotoUsecase(repository domains.PhotoRepository, userRepository domains.UserRepository, validate validation.Validation) domains.PhotoUsecase {
	return &PhotoUSecaseImpl{
		repository:     repository,
		UserRepository: userRepository,
		Validate:       validate,
	}
}

type PhotoUSecaseImpl struct {
	repository     domains.PhotoRepository
	UserRepository domains.UserRepository
	Validate       validation.Validation
}

func (usecase PhotoUSecaseImpl) CreatePhoto(request model.CreatePhotoRequest) (interface{}, string) {
	var response model.CreatePhotoResponse

	errorCode := make(chan string, 1)

	err := usecase.Validate.ValidateRequest(request)
	if err != nil {
		errorCode <- "400"
		return err, <-errorCode
	}

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
			PhotoUrl:  result.PhotoUrl,
			Caption:   result.Caption,
			Title:     result.Title,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
		errorCode <- "200"
	}

	return response, <-errorCode
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
func (usecase *PhotoUSecaseImpl) GetPhotoById(ctx *fiber.Ctx, id string) (*model.PhotoResponse, string) {
	responseCode := make(chan string, 1)
	result, err := usecase.repository.FindById(id)
	response := &model.PhotoResponse{}

	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)

	can, userId := helper.Can(email, "photo.list")
	if !can || userId != result.UserId {
		responseCode <- "403"
		return nil, <-responseCode
	}

	if result == nil && err == nil {
		responseCode <- "404"
		response = nil
		return nil, <-responseCode
	}

	if err != nil && result == nil {
		responseCode <- "500"
		response = nil
		return nil, <-responseCode
	}

	if err == nil && result != nil {

		responseCode <- "200"
		response = &model.PhotoResponse{
			Id:        result.Id,
			PhotoUrl:  result.PhotoUrl,
			Caption:   result.Caption,
			Title:     result.Title,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		}
	}
	return response, <-responseCode
}

// EditPhoto implements domains.PhotoUsecase
func (usecase *PhotoUSecaseImpl) EditPhoto(ctx *fiber.Ctx) (interface{}, string) {
	errCode := make(chan string, 1)
	response := &model.UpdatePhotoResponse{}
	var request model.CreatePhotoRequest
	ctx.BodyParser(&request)
	id := ctx.Params("id")
	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	result, err := usecase.repository.FindById(id)

	can, userId := helper.Can(email, "photo.edit")
	if !can || userId != result.UserId {
		errCode <- "403"
		return nil, <-errCode
	}

	errValidation := usecase.Validate.ValidateRequest(request)
	if errValidation != nil {
		errCode <- "400"
		return errValidation, <-errCode
	}

	if result == nil && err == nil {
		errCode <- "404"
		response = nil
		return response, <-errCode
	}

	if err != nil && result == nil {
		errCode <- "500"
		response = nil
		return response, <-errCode
	}

	result.PhotoUrl = request.PhotoUrl
	result.Caption = request.Caption
	result.Title = request.Title
	result.UpdatedAt = time.Now()

	err = usecase.repository.UpdatePhoto(id, *result)
	if err != nil {
		errCode <- "500"
		response = nil
		return response, <-errCode
	}
	errCode <- "200"
	return request, <-errCode
}

// DeletePhoto implements domains.PhotoUsecase
func (usecase *PhotoUSecaseImpl) DeletePhoto(ctx *fiber.Ctx) string {
	errCode := make(chan string, 1)
	id := ctx.Params("id")

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	result, err := usecase.repository.FindById(id)

	can, userId := helper.Can(email, "photo.delete")
	if !can || userId != result.UserId {
		errCode <- "403"
		return <-errCode
	}

	if result == nil && err == nil {
		errCode <- "404"
		return <-errCode
	}

	if err != nil && result == nil {
		errCode <- "500"
		return <-errCode
	}

	err = usecase.repository.DestroyPhoto(id)
	if err != nil {
		errCode <- "500"
		return <-errCode
	}

	errCode <- "200"
	return <-errCode
}
