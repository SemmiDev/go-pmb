package main

import (
	"github.com/gofiber/fiber/v2"
	"go-clean/config"
	"go-clean/controller"
	"go-clean/middleware"
	"go-clean/repository"
	"go-clean/service"
	"log"
	"os"
	"os/signal"
)

func main() {
	// setup configuration
	configuration := config.New()
	// setup database
	database := config.NewMongoDatabase(configuration)
	// setup repository
	registrationRepository := repository.NewRegistrationRepository(database)
	// setup service
	registrationService := service.NewRegistrationService(&registrationRepository)
	// Setup controller
	registrationController := controller.NewRegistrationController(&registrationService)
	// Setup fiber
	app := fiber.New()
	// Setup middleware
	middleware.FiberMiddleware(app)
	// Setup Routing
	registrationController.Route(app)

	// Run Server
	Run(app)
}

func Run(app *fiber.App) {
	// Create channel for idle connections.
	idleConsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := app.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConsClosed)
	}()

	// Run server.
	if err := app.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConsClosed
}
