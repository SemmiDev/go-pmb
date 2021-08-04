package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/controller"
	"github.com/SemmiDev/fiber-go-clean-arch/mailer"
	"github.com/SemmiDev/fiber-go-clean-arch/middleware"
	"github.com/SemmiDev/fiber-go-clean-arch/repository"
	"github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
)

func main() {

	// setup logger
	InitLogger()

	// setup configuration
	configuration := config.New()

	// setup database and token
	mongoDatabase := config.NewMongoDatabase(configuration)
	token := auth.NewToken()

	// setup repository
	registrationRepository := repository.NewRegistrationRepository(mongoDatabase)

	// setup mailer
	mailer := mailer.NewMail(config.NewMailDialer(configuration))

	// setup services
	redisService, err := config.NewRedisDB(configuration)
	if err != nil {
		zap.S().Error(err.Error())
		log.Fatal(err.Error())
	}
	registrationService := service.NewRegistrationService(&registrationRepository, &mailer)

	// Setup controller
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

// StartServer func for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
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
