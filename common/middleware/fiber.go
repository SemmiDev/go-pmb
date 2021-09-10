package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RegisterFiberMiddleware(app *fiber.App) {
	corsConfig := cors.ConfigDefault
	corsConfig.AllowCredentials = true
	app.Use(cors.New(corsConfig))

	loggerConfig := logger.ConfigDefault
	loggerConfig.TimeZone = "Asia/Jakarta"
	app.Use(logger.New(loggerConfig))

	app.Use(recover.New())
}
