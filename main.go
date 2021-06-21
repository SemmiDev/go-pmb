package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-clean/internal/app/controller"
	"go-clean/internal/app/repository"
	"go-clean/internal/app/service"
	"go-clean/internal/config"
	"go-clean/internal/exception"
	"os"
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
	// Setup Controller
	studentController := controller.NewStudentController(&studentService)
	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	// Setup Routing
	studentController.Route(app)

	// Start App
	err := app.Listen(os.Getenv("APP_PORT"))
	exception.PanicIfNeeded(err)
}
