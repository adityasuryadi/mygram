package model

import (
	"github.com/gofiber/fiber/v2"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Message string 		`json:"message"`
	Data   interface{} `json:"data"`
}

func SuccessResponse(ctx *fiber.Ctx,message string ,data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Message: message,
		Data:   data,
	}
	ctx.Status(fiber.StatusOK).JSON(response)
}

func NotFoundResponse(ctx *fiber.Ctx,message string , data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusNotFound,
		Status: "NOT_FOUND",
		Message: message,
		Data:   data,
	}
	ctx.Status(fiber.StatusNotFound).JSON(response)
}

func InternalServerErrorResponse(ctx *fiber.Ctx,message string , data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusInternalServerError,
		Status: "INTERNAL_SERVER_ERROR",
		Message: message,
		Data:   data,
	}
	ctx.Status(fiber.StatusInternalServerError).JSON(response)
}

func BadRequestResponse(ctx *fiber.Ctx,message string , data interface{}) {
	response := WebResponse{
		Code:   fiber.StatusBadRequest,
		Status: "BAD_REQUEST",
		Message: message,
		Data:   data,
	}
	ctx.Status(fiber.StatusBadRequest).JSON(response)
}

func ForbiddenResponse(ctx *fiber.Ctx,message string , data interface{}){
	response := WebResponse{
		Code:   fiber.StatusForbidden,
		Status: "UNAUTHORIZE",
		Message: message,
		Data:   data,
	}
	ctx.Status(fiber.StatusBadRequest).JSON(response)
}

func GetResponse(ctx *fiber.Ctx,errorCode string,message string,data interface{}){
	switch errorCode {
	case "200" :
		SuccessResponse(ctx,message,data)
	case "400":
		BadRequestResponse(ctx,message,data)
	case "403":
		ForbiddenResponse(ctx,message,data)
	case "404" :
		NotFoundResponse(ctx,message,data)
	case "500" :
		InternalServerErrorResponse(ctx,message,data)
	default:
		NotFoundResponse(ctx,message,data)
	}
}
