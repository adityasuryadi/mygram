package domains

import (
	entities "mygram/domains/entity"
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
)

// repository contract
type CommentRepository interface{
	StoreComment(comment entities.Comment) error
	GetComment()(comments []*entities.Comment,err error)
	FindCommentById(id string)(*entities.Comment,error)
	UpdateComment(id string,comment entities.Comment) error
	DestroyComment(id string) error
}

// usecase contract
type CommentUsecase interface {
	CreateComment(request *model.CreateCommentRequest) (interface{},string)
	GetAllComment()(comments []*model.CommentResponse,validation interface{},errCode string)
	GetCommentById(id string) (*model.CommentResponse,interface{},string)
	EditComment(ctx *fiber.Ctx)(string,interface{},*model.CommentResponse)
	DeleteComment(ctx *fiber.Ctx)(string,interface{},*model.CommentResponse)
}