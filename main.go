package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-clean/internal/app/controller"
	"go-clean/internal/app/repository"
	"go-clean/internal/app/service"
	"go-clean/internal/config"
	"go-clean/internal/exception"
)

func main() {
	// setup configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	// setup repository
	counterRepo := repository.NewCounterRepo(database)
	studentRepository := repository.NewStudentRepository(database, counterRepo)
	// setup service
	studentService := service.NewStudentService(&studentRepository)
	// Setup Controller
	studentController := controller.NewStudentController(&studentService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	studentController.Route(app)

	// Start App
	err := app.Listen(":9090")
	exception.PanicIfNeeded(err)
}