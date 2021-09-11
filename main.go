package main

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/src/common/config"
	"github.com/SemmiDev/go-pmb/src/registrant"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"log"
	"net/http"
)

func main() {
	// load app config.
	config.Load()

	// creates a new Fiber instance.
	app := fiber.New()

	// register middlewares.
	corsConfig := cors.ConfigDefault
	corsConfig.AllowCredentials = true
	app.Use(cors.New(corsConfig))
	app.Use(recover.New())

	// register health check endpoint.
	app.Get("/api/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON("i'm healthy")
	})

	registrantServer, err := registrant.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	registrantServer.Mount(app)

	fmt.Printf("running on http://localhost%v", config.AppPort)
	log.Fatal(app.Listen(config.AppPort))
}
