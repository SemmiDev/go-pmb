package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-clean/config"
	"go-clean/controller"
	"go-clean/exception"
	"go-clean/repository"
	"go-clean/service"
)

func main() {
	// setup configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	// setup repository
	studentRepository := repository.NewStudentRepository(database)
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