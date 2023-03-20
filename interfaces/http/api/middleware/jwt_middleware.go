package middleware

import (
	"mygram/domains/model"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func Verify() fiber.Handler {
	config := jwtware.Config{
		ErrorHandler: jwtError,
		SigningKey:   []byte("secret"),
		SuccessHandler: func(c *fiber.Ctx) error {
			return c.Next()
		},
	}
	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// Return status 401 and failed authentication error.
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "BAD_REQUEST",
			Message:   err.Error(),
			Data: nil,
		})
	}

	// Return status 401 and failed authentication error.
	return c.Status(fiber.StatusUnauthorized).JSON(model.WebResponse{
		Code:   fiber.StatusUnauthorized,
		Status: "UNAUTHORIZE",
		Message: "UNAUTHORIZE",
		Data:   nil,
	})
}
