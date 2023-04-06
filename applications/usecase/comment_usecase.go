package usecase

import (
	"log"
	"mygram/domains"
	entities "mygram/domains/entity"
	"mygram/domains/model"
	"mygram/infrastructures/validation"
	"time"

	"github.com/google/uuid"
)

func NewCommmentUsecase(commentRepository domains.CommentRepository, userRepository domains.UserRepository, validate validation.Validation) domains.CommentUsecase {
	return &CommentUsecaseImpl{
		CommentRepo: commentRepository,
		UserRepo:    userRepository,
		Validate:    validate,
	}
}

type CommentUsecaseImpl struct {
	CommentRepo domains.CommentRepository
	UserRepo    domains.UserRepository
	Validate    validation.Validation
}

// DeleteComment implements domains.CommentUsecase
func (commentUsecase *CommentUsecaseImpl) DeleteComment(id string) (string, interface{}, *model.CommentResponse) {
	errCode := make(chan string,1)
	result, err := commentUsecase.CommentRepo.FindCommentById(id)

	if result == nil && err == nil {
		errCode <- "404"
		return <-errCode,nil,nil
	}

	if err != nil && result == nil {
		errCode <- "500"
		return <-errCode,nil,nil
	}

	err = commentUsecase.CommentRepo.DestroyComment(id)
	if err != nil {
		errCode <- "500"
		return <-errCode,nil,nil
	}
	
	errCode <- "200"
	return <-errCode,nil,nil
}

// EditComment implements domains.CommentUsecase
func (commentUsecase *CommentUsecaseImpl) EditComment(id string, request model.UpdateCommentRequest) (string, interface{}, *model.CommentResponse) {
	errCode := make(chan string, 1)
	var response = &model.CommentResponse{}

	err := commentUsecase.Validate.ValidateRequest(request)
	if err != nil {
		errCode <- "400"
		return <-errCode, err, nil
	}

	result, err := commentUsecase.CommentRepo.FindCommentById(id)

	if result == nil && err == nil {
		errCode <- "404"
		return <-errCode, nil, nil
	}

	if err != nil && result == nil {
		errCode <- "500"
		return <-errCode, nil, nil
	}

	result.Message = request.Message
	result.PhotoId = request.PhotoId
	result.UpdatedAt = time.Now()

	err = commentUsecase.CommentRepo.UpdateComment(id, *result)
	if err != nil {
		errCode <- "500"
		return <-errCode, nil, nil
	}
	errCode <- "200"
	return <-errCode, nil, response
}

// GetCommentById implements domains.CommentUsecase
func (commentUsecase *CommentUsecaseImpl) GetCommentById(id string) (*model.CommentResponse, interface{}, string) {
	errCode := make(chan string, 1)
	result, err := commentUsecase.CommentRepo.FindCommentById(id)

	log.Print(result, err)
	if result == nil && err == nil {
		errCode <- "404"
		return nil, nil, <-errCode
	}

	if err != nil && result == nil {
		errCode <- "500"
		return nil, nil, <-errCode
	}

	response := &model.CommentResponse{
		Id:        result.Id,
		PhotoId:   result.PhotoId,
		Message:   result.Message,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	errCode <- "200"
	return response, nil, <-errCode
}

// GetAllComment implements domains.CommentUsecase
func (commentUsecase *CommentUsecaseImpl) GetAllComment() (comments []*model.CommentResponse, validation interface{}, errCode string) {
	errorCode := make(chan string, 1)
	results, err := commentUsecase.CommentRepo.GetComment()
	for _, v := range results {
		comments = append(comments, &model.CommentResponse{
			Id:        v.Id,
			PhotoId:   v.PhotoId,
			Message:   v.Message,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	if err != nil {
		errorCode <- "500"
		return nil, nil, <-errorCode
	}
	errorCode <- "200"
	return comments, nil, <-errorCode
}

func (commentUsecase *CommentUsecaseImpl) CreateComment(request *model.CreateCommentRequest) (interface{}, string) {
	var comment = entities.Comment{
		Message: request.Message,
		PhotoId: request.PhotoId,
		UserId:  uuid.New(),
	}

	errorCode := make(chan string, 1)

	err := commentUsecase.Validate.ValidateRequest(request)
	if err != nil {
		errorCode <- "400"
		return err, <-errorCode
	}

	err = commentUsecase.CommentRepo.StoreComment(comment)
	if err != nil {
		errorCode <- "500"
		return nil, <-errorCode
	}

	errorCode <- "200"
	return nil, <-errorCode
}
