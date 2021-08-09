package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/setup"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Setup Fiber
	app := fiber.New()

	// Setup App
	setup.App(app)

	// Start Server
	// setup.StartServerWithGracefulShutdown(app)
	setup.StartServer(app)
}
