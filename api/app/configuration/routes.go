package configuration

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/app/routes"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")

	routes.AddHealthRouter(v1)
	routes.AddRegistrationRouter(v1)
	routes.AddAuthRouter(v1)
}
