package main

import (
	"github.com/gofiber/fiber/v2"
	"go-clean/internal/app/controller"
	"go-clean/internal/app/repository"
	"go-clean/internal/app/service"
	"go-clean/internal/config"
	"go-clean/internal/middleware"
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
	studentRepository := repository.NewStudentRepository(database)
	// setup service
	studentService := service.NewStudentService(&studentRepository)
	// Setup controller
	studentController := controller.NewStudentController(&studentService)
	// Setup fiber
	app := fiber.New(config.NewFiberConfig())
	// Setup middleware
	middleware.FiberMiddleware(app)
	// Setup Routing
	studentController.Route(app)

	// Start server
	StartServer(app)
}

func StartServer(app *fiber.App) {
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
	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConsClosed
}
