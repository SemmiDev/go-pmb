package controllers

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/requests"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/responses"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/gofiber/fiber/v2"
)

type RegistrationController struct {
	Service interfaces.IRegistrationService
}

func NewRegistrationController(u interfaces.IRegistrationService) *RegistrationController {
	return &RegistrationController{Service: u}
}

func (rc *RegistrationController) Register(c *fiber.Ctx) error {
	request := new(requests.Register)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.UnprocessableEntity(err))
	}

	err := request.IsValid()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadRequest(err))
	}

	response := rc.Service.Register(request)
	return c.Status(response.Meta.Code).JSON(response)
}

func (rc *RegistrationController) UpdatePaymentStatus(c *fiber.Ctx) error {
	request := new(requests.UpdatePaymentStatus)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(responses.UnprocessableEntity(err))
	}

	err := request.IsValid()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.BadRequest(err))
	}

	response := rc.Service.UpdatePaymentStatus(request)
	return c.Status(response.Meta.Code).JSON(response)
}
