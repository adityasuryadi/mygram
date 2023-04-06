package handler

import (
	"mygram/domains"
	"mygram/domains/model"
	"mygram/interfaces/http/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewCommentHandler(commentUsecase domains.CommentUsecase) CommentHandler{ 
	return CommentHandler{
		CommentUsecase: commentUsecase,
	}
}

type CommentHandler struct {
	CommentUsecase domains.CommentUsecase
}

func (handler CommentHandler) Route(app *fiber.App){
	comment:=app.Group("comment",middleware.Verify())
	comment.Post("/",handler.PostComment)
	comment.Get("/",handler.GetAllComment)
	comment.Get("/:id",handler.GetOneComment)
	comment.Put("/:id",handler.UpdateComment)
}

func (handler CommentHandler) PostComment(ctx *fiber.Ctx) error{
	var request model.CreateCommentRequest
	ctx.BodyParser(&request)
	result,errCode := handler.CommentUsecase.CreateComment(&request)

	if errCode == "400" {
		model.BadRequestResponse(ctx,"CLIENT ERROR",result)
		return nil
	}

	if errCode == "500" {
		model.InternalServerErrorResponse(ctx,"CLIENT ERROR",nil)
		return nil
	}

	if errCode == "200" {
		model.SuccessResponse(ctx,"CLIENT ERROR",nil)
		return nil
	}

	return nil
}

func (handler CommentHandler) GetAllComment(ctx *fiber.Ctx) error {
	result,_,errCode := handler.CommentUsecase.GetAllComment()
	model.GetResponse(ctx,errCode,"",result)
	return nil
}

func (handler CommentHandler) GetOneComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result,_,errCode := handler.CommentUsecase.GetCommentById(id)
	model.GetResponse(ctx,errCode,"",result)
	return nil 
}

func (handler CommentHandler) UpdateComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result,validation,errCode:=handler.CommentUsecase.GetCommentById(id)
	if errCode == "400" {
		model.GetResponse(ctx,errCode,"",validation)
	}
		model.GetResponse(ctx,errCode,"",result)
		return nil	
}

func (handler CommentHandler) DeleteComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	errCode,_,result:=handler.CommentUsecase.DeleteComment(id)
	model.GetResponse(ctx,errCode,"",result)
	return nil
}