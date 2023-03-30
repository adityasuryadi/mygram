package fiberconfig

import (
	"mygram/commons/exceptions"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exceptions.ErrorHandler,
	}
}