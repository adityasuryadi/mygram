package usecase

import (
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/helper"
	"mygram/infrastructures/security"
	"mygram/infrastructures/validation"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewSocialmediaUsecase(socialmediaRepo domains.SocialmediaRepository,userRepo domains.UserRepository,validate validation.Validation ) domains.SocialmediaUsecase {
	return &SocialmediaUsecaseImpl{
		socialmediaRepo:  socialmediaRepo,
		userRepo: userRepo,
		validate: validate,
	}
}

type SocialmediaUsecaseImpl struct {
	socialmediaRepo domains.SocialmediaRepository
	userRepo domains.UserRepository
	validate validation.Validation
}

func (socialmediaUsecase *SocialmediaUsecaseImpl) CreateSocialmedia(request model.CreateSocialmediaRequest) (string,interface{},*model.SocialmediaResponse){
	
	responseCode := make(chan string, 1)
	err:=socialmediaUsecase.validate.ValidateRequest(request)
	if err != nil {
		responseCode <- "400"
		return <-responseCode,err,nil
	}

	result,err := socialmediaUsecase.userRepo.GetUserByEmail(request.Email)

	if err != nil {
		responseCode <- "403"
		return <-responseCode,nil,nil
	}

	socialmedia := entities.Socialmedia{
		Name: request.Name,
		SocialMediaUrl: request.SocialmediaUrl,
		UserId: result.Id,
	}

	err = socialmediaUsecase.socialmediaRepo.Insert(socialmedia)
	
	if err != nil {
		responseCode <- "500"
		return <-responseCode,nil,nil
	}
	
	responseCode <- "200"
	return <-responseCode,nil,nil
}

func (socialmediaUsecase *SocialmediaUsecaseImpl) ListSocialmedia()(string,interface{},[]*model.SocialmediaResponse){
	responseCode := make(chan string,1)
	var responses []*model.SocialmediaResponse
	result,err := socialmediaUsecase.socialmediaRepo.GetAll()
	if err != nil {
		responseCode <- "500"
		return <-responseCode,nil,nil
	}

	responseCode <- "200"

	for _, v := range result {
		responses = append(responses, &model.SocialmediaResponse{
			Id: v.Id,
			Name: v.Name,
			SocialmediaUrl: v.SocialMediaUrl,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	
	return <-responseCode,nil,responses
}

func (socialmediaUsecase *SocialmediaUsecaseImpl) FindSocialmediaById(id string) (string,interface{},*model.SocialmediaResponse){
	responseCode := make(chan string,1)
	result,err := socialmediaUsecase.socialmediaRepo.FindById(id)
	var response = &model.SocialmediaResponse{}

	if result == nil && err == nil {
		responseCode <- "404"
		response = nil
	}

	if err != nil && result == nil {
		responseCode <- "500"
		response = nil
	}

	if err == nil && result != nil {

		responseCode <- "200"
		response = &model.SocialmediaResponse{
			Id:             result.Id,
			Name:           result.Name,
			SocialmediaUrl: result.SocialMediaUrl,
			CreatedAt:      result.CreatedAt,
			UpdatedAt:      result.UpdatedAt,
		}
	}
	return <-responseCode,nil,response
}

func (socialmediaUsecase *SocialmediaUsecaseImpl) EditSocialmedia(ctx *fiber.Ctx)(string,interface{},*model.SocialmediaResponse){
	responseCode := make(chan string,1)
	var request model.CreateSocialmediaRequest
	id:=ctx.Params("id")

	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)
	

	socialmedia,err := socialmediaUsecase.socialmediaRepo.FindById(id)
	can,userId:=helper.Can(email,"socialmedia.edit")
	if !can || userId != socialmedia.UserId {
		responseCode <- "403"
		return <-responseCode,nil,nil
	}
	
	if socialmedia == nil && err == nil {
		responseCode <- "404"
		return <-responseCode,nil,nil
	}

	validationErr := socialmediaUsecase.validate.ValidateRequest(request)
	if validationErr != nil {
		responseCode<-"400"
		return <-responseCode,validationErr,nil
	}

	socialmedia.Name = request.Name
	socialmedia.SocialMediaUrl = request.SocialmediaUrl
	entities,err := socialmediaUsecase.socialmediaRepo.Update(socialmedia)
	
	response:= model.SocialmediaResponse{
		Id:entities.Id,
		Name: entities.Name,
		SocialmediaUrl: entities.SocialMediaUrl,
		CreatedAt: entities.CreatedAt,
		UpdatedAt: entities.UpdatedAt,
	}
	responseCode<-"200"
	return <-responseCode,nil,&response
}

func (socialmediaUsecase *SocialmediaUsecaseImpl) DeleteSocialmedia(ctx *fiber.Ctx) (string,interface{},*model.SocialmediaResponse){
	responseCode := make(chan string,1)
	id := ctx.Params("id")
	
	claims := security.DecodeToken(ctx.Locals("user").(*jwt.Token))
	email := claims["email"].(string)

	socialmedia,err := socialmediaUsecase.socialmediaRepo.FindById(id)

	can,userId:=helper.Can(email,"socialmedia.delete")
	if !can || userId != socialmedia.UserId {
		responseCode <- "403"
		return <-responseCode,nil,nil
	}

	if socialmedia == nil && err == nil {
		responseCode <- "404"
		return <-responseCode,nil,nil
	}
	err = socialmediaUsecase.socialmediaRepo.Delete(id)
	if err != nil {
		responseCode <- "500"
	}else{
		responseCode<-"200"	
	}	
	return <-responseCode,nil,nil
}
