package tests

import (
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/controller"
	"github.com/SemmiDev/fiber-go-clean-arch/middleware"
	"github.com/SemmiDev/fiber-go-clean-arch/repository"
	"github.com/SemmiDev/fiber-go-clean-arch/tests/fakeservice"
	"github.com/gofiber/fiber/v2"
)

func createTestApp() *fiber.App {
	var app = fiber.New()
	middleware.FiberMiddleware(app)
	registrationController.Route(app)
	return app
}

// note: make sure asynq email server start first
var configuration = config.New("../.env")
var database = config.NewMongoDatabase(configuration)
var registrationRepository = repository.NewRegistrationRepository(database)
var registrationService = fakeservice.NewRegistrationService(&registrationRepository)

var token = auth.NewToken()
var redisService, err = config.NewRedisDB(configuration)
var registrationController = controller.NewRegistrationController(&registrationService, redisService.Auth, token)

var app = createTestApp()
