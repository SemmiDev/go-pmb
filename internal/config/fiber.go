package config

import (
	"github.com/gofiber/fiber/v2"
	exception2 "go-clean/internal/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception2.ErrorHandler,
	}
}
