package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/controller"
	"github.com/SemmiDev/fiber-go-clean-arch/logger"
	"github.com/SemmiDev/fiber-go-clean-arch/middleware"
	"github.com/SemmiDev/fiber-go-clean-arch/payment"
	"github.com/SemmiDev/fiber-go-clean-arch/repository"
	"github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Setup fiber
	app := fiber.New()

	// Setup App
	SetupApp(app)

	// StartServer
	StartServer(app)
}

// SetupApp setup many stuff.
func SetupApp(app *fiber.App) {
	// Setup middleware
	middleware.FiberMiddleware(app)

	// setup logger
	logger.SetupLogger()

	// setup configuration
	configuration := config.New()

	// setup database
	mongoDatabase := config.NewMongoDatabase(configuration)

	// setup token
	token := auth.NewToken()

	// setup repository
	registrationRepository := repository.NewRegistrationRepository(mongoDatabase)

	// setup message broker
	amqpServerURL := configuration.Get("AMQP_SERVER_URL")
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	panicIfNeeded(err)
	defer connectRabbitMQ.Close()
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	panicIfNeeded(err)
	defer channelRabbitMQ.Close()

	// setup queue name for rabbitMQ
	_, err = channelRabbitMQ.QueueDeclare(
		"QueueEmailServiceRegistration", // queue name
		true,                            // durable
		false,                           // auto delete
		false,                           // exclusive
		false,                           // no wait
		nil,                             // arguments
	)
	panicIfNeeded(err)

	// setup redis
	redisService, err := config.NewRedisDB(configuration)
	panicIfNeeded(err)

	// setup midtrans client
	midtransClient := config.NewMidtransClient(configuration)

	// setup midtrans service
	midtransService := payment.NewService(midtransClient)

	// setup registration service
	registrationService := service.NewRegistrationService(&registrationRepository, channelRabbitMQ, midtransService)

	// Setup controllers
	registrationController := controller.NewRegistrationController(&registrationService)
	authController := controller.NewAuthController(
		&registrationService,
		redisService.Auth,
		token,
	)

	// Setup Routing
	registrationController.Route(app)
	authController.Route(app)
}

// panicIfNeeded do panic if err != nil.
func panicIfNeeded(err error) {
	if err != nil {
		panic(err)
	}
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
