package main

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/internal/common/config"
	"github.com/SemmiDev/go-pmb/internal/registrant/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"net/http"
)

func main() {
	config.Load()
	app := fiber.New()

	// middlewares
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("OK")
	})
	registrant := server.NewRegistrantServer()
	registrant.Mount(app)

	fmt.Printf("🚀 API is running on http://localhost%v", config.AppPort)
	log.Fatal(app.Listen(config.AppPort))
}
