package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/controller"
	"github.com/SemmiDev/fiber-go-clean-arch/middleware"
	"github.com/SemmiDev/fiber-go-clean-arch/repository"
	"github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
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
	asyncRedisConnection := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_DSN"), // Redis server address
	}
	client := asynq.NewClient(asyncRedisConnection)
	defer client.Close()

	registrationService := service.NewRegistrationService(&registrationRepository, client)

	// Setup controller
	token := auth.NewToken()
	redisService, err := config.NewRedisDB(
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
		os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}

	registrationController := controller.NewRegistrationController(
		&registrationService,
		redisService.Auth,
		token,
	)
	// Setup fiber
	app := fiber.New()
	// Setup middleware
	middleware.FiberMiddleware(app)
	// Setup Routing
	registrationController.Route(app)

	// StartServer
	StartServer(app)
}

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(a *fiber.App) {
	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // Catch OS signals.
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := a.Shutdown(); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := a.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}
