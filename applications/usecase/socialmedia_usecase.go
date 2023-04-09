package usecase

import (
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/validation"
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
