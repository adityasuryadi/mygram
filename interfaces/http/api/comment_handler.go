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
	comment.Delete("/:id",handler.DeleteComment)
}

// PostComment godoc
// @Summary Create Comment
// @Description Create Comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body model.CreateCommentRequest true "Create New Comment"
// @Success 200 {object} model.WebResponse{data=model.CommentResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /comment [post]
// @Security BearerAuth
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

// List Comment
// @Summary List Comment
// @Description List Comment
// @Tags comment
// @Accept json
// @Produce json
// @Success 200 {object} model.WebResponse{data=[]model.CommentResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /comment [get]
// @Security BearerAuth
func (handler CommentHandler) GetAllComment(ctx *fiber.Ctx) error {
	result,_,errCode := handler.CommentUsecase.GetAllComment()
	model.GetResponse(ctx,errCode,"",result)
	return nil
}


// Find Comment
// @Summary Find Comment
// @Description Find Comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment Id"
// @Success 200 {object} model.WebResponse{data=model.CommentResponse}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /comment/{id} [get]
// @Security BearerAuth
func (handler CommentHandler) GetOneComment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	result,_,errCode := handler.CommentUsecase.GetCommentById(id)
	model.GetResponse(ctx,errCode,"",result)
	return nil 
}


// Edit Comment
// @Summary Edit Comment
// @Description Edit Comment
// @Tags comment
// @Accept json
// @Produce json
// @Param comment body model.UpdateCommentRequest true "Update Comment"
// @Param id path string true "Comment Id"
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /comment/{id} [put]
// @Security BearerAuth
func (handler CommentHandler) UpdateComment(ctx *fiber.Ctx) error {
	responseCode,validation,result:=handler.CommentUsecase.EditComment(ctx)
	if responseCode == "400" {
		model.GetResponse(ctx,responseCode,"",validation)
	}
		model.GetResponse(ctx,responseCode,"",result)
		return nil	
}


// Delete Comment
// @Summary Delete Comment
// @Description Delete Comment
// @Tags comment
// @Accept json
// @Produce json
// @Param id path string true "Comment Id"
// @Success 200 {object} model.WebResponse{}
// @Failure 400 {object} model.WebResponse{}
// @Failure 500 {object} model.WebResponse{}
// @Router /comment/{id} [delete]
// @Security BearerAuth
func (handler CommentHandler) DeleteComment(ctx *fiber.Ctx) error {
	errCode,_,result:=handler.CommentUsecase.DeleteComment(ctx)
	model.GetResponse(ctx,errCode,"",result)
	return nil
}