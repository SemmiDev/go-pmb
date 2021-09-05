package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RegisterFiberMiddleware(app *fiber.App) {
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())
}
