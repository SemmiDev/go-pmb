package main

import (
	auth2 "github.com/SemmiDev/fiber-go-clean-arch/auth"
	config2 "github.com/SemmiDev/fiber-go-clean-arch/config"
	controller2 "github.com/SemmiDev/fiber-go-clean-arch/controller"
	mailer2 "github.com/SemmiDev/fiber-go-clean-arch/mailer"
	"github.com/SemmiDev/fiber-go-clean-arch/middleware"
	repository2 "github.com/SemmiDev/fiber-go-clean-arch/repository"
	service2 "github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

func main() {
	// setup configuration
	configuration := config2.New()

	// setup database and token
	mongoDatabase := config2.NewMongoDatabase(configuration)
	token := auth2.NewToken()

	// setup repository
	registrationRepository := repository2.NewRegistrationRepository(mongoDatabase)

	// setup mailer
	mailer := mailer2.NewMail(config2.NewMailDialer(configuration))

	// setup services
	redisService, err := config2.NewRedisDB(configuration)
	if err != nil {
		log.Fatal(err.Error())
	}
	registrationService := service2.NewRegistrationService(&registrationRepository, &mailer)

	// Setup controller
	registrationController := controller2.NewRegistrationController(
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
