package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-clean/config"
	"go-clean/repository"
	"go-clean/service"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	studentController.Route(app)
	return app
}

var configuration = config.New("../.env.test")

var database = config.NewMongoDatabase(configuration)
var studentRepository = repository.NewStudentRepository(database)
var studentService = service.NewStudentService(&studentRepository)
var studentController = NewStudentController(&studentService)
var app = createTestApp()