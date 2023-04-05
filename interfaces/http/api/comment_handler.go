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
