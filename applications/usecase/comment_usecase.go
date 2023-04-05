package usecase

import (
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/validation"

	"github.com/google/uuid"
)

func NewCommmentUsecase(commentRepository domains.CommentRepository,userRepository domains.UserRepository,validate validation.Validation) domains.CommentUsecase{
	return &CommentUsecaseImpl{
		CommentRepo: commentRepository,
		UserRepo: userRepository,
		Validate: validate,
	}
}

type CommentUsecaseImpl struct{
	CommentRepo domains.CommentRepository
	UserRepo domains.UserRepository
	Validate validation.Validation
}

func (commentUsecase *CommentUsecaseImpl) CreateComment(request *model.CreateCommentRequest) (interface{},string){
	var comment = entities.Comment{
		Message: request.Message,
		PhotoId: request.PhotoId,
		UserId: uuid.New(),
	}

	errorCode := make(chan string,1)

	err := commentUsecase.Validate.ValidateRequest(request)
	if err != nil {
		errorCode <- "400"
		return err,<-errorCode
	}

	err = commentUsecase.CommentRepo.StoreComment(comment)
	if err != nil {
		errorCode <- "500"
		return nil,<-errorCode	
	}

	errorCode <- "200"
	return nil,<-errorCode
}