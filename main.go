package main

import (
	"fmt"
	"github.com/SemmiDev/go-pmb/api/registrant"
	"github.com/SemmiDev/go-pmb/common/config"
	"github.com/SemmiDev/go-pmb/common/middleware"
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

	registrantServer, err := registrant.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	registrantServer.Mount(app)

	fmt.Printf("running on http://localhost%v", config.AppPort)
	log.Fatal(app.Listen(config.AppPort))
}
