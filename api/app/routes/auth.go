package routes

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/ioc"
	"github.com/gofiber/fiber/v2"
)

func AddAuthRouter(router fiber.Router) {
	router.Post("/auth/login", ioc.AuthController.Login)
	router.Post("/auth/logout", ioc.AuthController.Logout)
	router.Post("/auth/refresh", ioc.AuthController.Refresh)
}
