package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/controller"
	"github.com/SemmiDev/fiber-go-clean-arch/middleware"
	"github.com/SemmiDev/fiber-go-clean-arch/repository"
	"github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
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

	// setup message broker
	amqpServerURL := configuration.Get("AMQP_SERVER_URL")

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	panicIfNeeded(err)
	defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	panicIfNeeded(err)
	defer channelRabbitMQ.Close()

	// setup queue name
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

	// setup service
	registrationService := service.NewRegistrationService(&registrationRepository, channelRabbitMQ)

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

// panicIfNeeded panic if err != nil
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
