package main

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/internal/common/config"
	"github.com/SemmiDev/go-pmb/internal/common/middleware"
	"github.com/SemmiDev/go-pmb/internal/registrant/server"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func main() {
	config.Load()

	app := fiber.New()
	middleware.RegisterFiberMiddleware(app)

	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("i'm healthy")
	})
	registrant := server.NewRegistrantServer()
	registrant.Mount(app)

	fmt.Printf("running on http://localhost%v", config.AppPort)
	log.Fatal(app.Listen(config.AppPort))
}
