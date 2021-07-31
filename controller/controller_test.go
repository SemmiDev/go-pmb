package controller

import (
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/middleware"
	"github.com/SemmiDev/fiber-go-clean-arch/repository"
	"github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/gofiber/fiber/v2"
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
