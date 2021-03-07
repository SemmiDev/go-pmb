package config

import (
	"github.com/gofiber/fiber/v2"
	"go-clean/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
