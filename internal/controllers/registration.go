package controllers

import (
	"github.com/SemmiDev/fiber-go-clean-arch/internal/models"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/services"
	"github.com/gofiber/fiber/v2"
)

type RegistrationController struct {
	RegistrationService services.RegistrationService
}

func NewRegistrationController(registrationService *services.RegistrationService) RegistrationController {
	return RegistrationController{
		RegistrationService: *registrationService,
	}
}

func (c *RegistrationController) Route(app *fiber.App) {
	v1 := app.Group("/api/v1")

	v1.Post("/register", c.Register)
	v1.Put("/register/payment/status", c.UpdatePaymentStatus)
}

func (c *RegistrationController) Register(ctx *fiber.Ctx) error {
	var request models.RegistrationRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(models.APIResponse("Cannot unmarshal body", fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
	}

	errs := request.Validate()
	if errs != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(models.APIResponse(errs, fiber.StatusBadRequest, "Bad Request", nil))
	}

	response, err := c.RegistrationService.Register(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(models.APIResponse(err.Error(), fiber.StatusBadRequest, "Bad Request", nil))
	}

	return ctx.Status(fiber.StatusCreated).
		JSON(models.APIResponse(nil, fiber.StatusCreated, "Created", response))
}

func (c *RegistrationController) UpdatePaymentStatus(ctx *fiber.Ctx) error {
	var request models.UpdatePaymentStatusRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(models.APIResponse("Cannot unmarshal body", fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
	}

	errs := request.Validate()
	if errs != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(models.APIResponse(errs, fiber.StatusBadRequest, "Unprocessable Entity", nil))
	}

	status, err := c.RegistrationService.UpdatePaymentStatus(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(models.APIResponse(err.Error(), fiber.StatusBadRequest, "Bad Request", nil))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(models.APIResponse("Updated", fiber.StatusOK, "Ok", map[string]string{
			"payment_status": status,
		}))
}
