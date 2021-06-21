package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
		// add recoverer for panic
		recover.New(),
	)
}
