package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/internal/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/config"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/controllers"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/helper"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/middleware"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/repositories"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Setup fiber.
	app := fiber.New()

	// Setup app.
	SetupApp(app)

	// Start server.
	// app.StartServerWithGracefulShutdown(app)
	StartServer(app)
}

// SetupApp app many stuff.
func SetupApp(app *fiber.App) {
	// Setup middleware.
	middleware.FiberMiddleware(app)

	// setup logger.
	helper.SetupLogger()

	// setup configuration.
	configuration := config.New()

	// setup database.
	mongoDatabase := config.NewMongoDatabase(configuration)

	// setup token.
	token := auth.NewToken()

	// setup repositories.
	registrationRepository := repositories.NewRegistrationRepository(mongoDatabase)

	// setup message broker.
	amqpServerURL := configuration.Get("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	helper.PanicIfNeeded(err)
	defer connectRabbitMQ.Close()
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	helper.PanicIfNeeded(err)
	defer channelRabbitMQ.Close()

	// setup queue name for rabbitMQ.
	_, err = channelRabbitMQ.QueueDeclare(
		"QueueEmailServiceRegistration", // queue name
		true,                            // durable
		false,                           // auto delete
		false,                           // exclusive
		false,                           // no wait
		nil,                             // arguments
	)
	helper.PanicIfNeeded(err)

	// setup redis.
	redisService, err := config.NewRedisDB(configuration)
	helper.PanicIfNeeded(err)

	// setup midtrans client.
	midtransClient := config.NewMidtransClient(configuration)

	// setup midtrans services.
	midtransService := services.NewService(midtransClient)

	// setup registration services.
	registrationService := services.NewRegistrationService(&registrationRepository, channelRabbitMQ, midtransService)

	// Setup controllers.
	registrationController := controllers.NewRegistrationController(&registrationService)
	authController := controllers.NewAuthController(
		&registrationService,
		redisService.Auth,
		token,
	)

	// Setup routes.
	registrationController.Route(app)
	authController.Route(app)
}

// StartServer  for starting a simple server.
func StartServer(a *fiber.App) {
	// Run server.
	if err := a.Listen(os.Getenv("APP_PORT")); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}
}

// StartServerWithGracefulShutdown for starting server with a graceful shutdown.
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
