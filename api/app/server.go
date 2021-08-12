package app

import (
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
)

func Run(a *fiber.App) {
	fmt.Printf("🚀 API is running on http://localhost%v", environments.AppPort)
	log.Fatal(a.Listen(environments.AppPort))
}

// RunServerWithGracefulShutdown for starting server with a graceful shutdown.
func RunServerWithGracefulShutdown(a *fiber.App) {
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
	if err := a.Listen(environments.AppPort); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
