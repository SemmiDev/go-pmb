package routes

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/ioc"
	"github.com/gofiber/fiber/v2"
)

func AddRegistrationRouter(router fiber.Router) {
	router.Post("/register", ioc.RegistrationController.Register)
	router.Put("/register/paymentStatus", ioc.RegistrationController.UpdatePaymentStatus)
}
