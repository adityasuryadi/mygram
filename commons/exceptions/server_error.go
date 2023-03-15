package exceptions

import (
	"time"

	log "mygram/infrastructures/log"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	errData := logrus.Fields{
		"at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": string(ctx.Request().Header.Method()),
		"uri":    ctx.Request().URI().String(),
		"ip":     ctx.IP(),
		"error":  err.Error(),
	}
	log.WriteLog(errData)
	return ctx.Status(fiber.StatusInternalServerError).JSON(WebResponse{
		Code: 500,
		Status: "INTERNAL_SERVER_ERROR",
		Data: nil,
	})
}