package configuration

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/database"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/ioc"
	"github.com/SemmiDev/fiber-go-clean-arch/api/presentation/middlewares"
	"github.com/gofiber/fiber/v2"
)

func App() *fiber.App {
	app := fiber.New()

	// utils.SetupLogger()

	middlewares.FiberMiddleware(app)

	environments.New()

	mongoConnection := database.NewMongoConnection()
	redisConnection := database.NewRedisConnection()

	ioc.SetupDependencyInjection(mongoConnection, redisConnection)
	SetupRoutes(app)

	return app
}
