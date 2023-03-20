package model

import "github.com/gofiber/fiber/v2"

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
