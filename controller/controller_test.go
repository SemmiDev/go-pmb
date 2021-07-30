package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-clean/config"
	"go-clean/middleware"
	"go-clean/repository"
	"go-clean/service"
)

func createTestApp() *fiber.App {
	var app = fiber.New()
	middleware.FiberMiddleware(app)
	registrationController.Route(app)
	return app
}

var configuration = config.New("../.env")
var database = config.NewMongoDatabase(configuration)
var registrationRepository = repository.NewRegistrationRepository(database)
var registrationService = service.NewRegistrationService(&registrationRepository)
var registrationController = NewRegistrationController(&registrationService)

var app = createTestApp()
