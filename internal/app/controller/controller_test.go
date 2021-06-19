package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-clean/internal/app/repository"
	"go-clean/internal/app/service"
	"go-clean/internal/config"
)

func createTestApp() *fiber.App {
	var app = fiber.New(config.NewFiberConfig())
	app.Use(recover.New())
	studentController.Route(app)
	return app
}

var configuration = config.New("../../../.env")

var database = config.NewMongoDatabase(configuration)
var counterRepository = repository.NewCounterRepo(database)
var studentRepository = repository.NewStudentRepository(database, counterRepository)
var studentService = service.NewStudentService(&studentRepository)
var studentController = NewStudentController(&studentService)
var app = createTestApp()