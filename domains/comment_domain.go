package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"
)

// repository contract
type CommentRepository interface{
	StoreComment(comment entities.Comment) error
}

// usecase contract
type CommentUsecase interface {
	CreateComment(request *model.CreateCommentRequest) (interface{},string)
}