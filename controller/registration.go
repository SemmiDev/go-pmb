package controller

import (
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/service"
	"github.com/gofiber/fiber/v2"
)

type RegistrationController struct {
	RegistrationService service.RegistrationService
}

func NewRegistrationController(registrationService *service.RegistrationService) RegistrationController {
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
	var request model.RegistrationRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(model.APIResponse("Cannot unmarshal body", fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
	}

	errs := request.Validate()
	if errs != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.APIResponse(errs, fiber.StatusBadRequest, "Bad Request", nil))
	}

	response, err := c.RegistrationService.Register(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.APIResponse(err.Error(), fiber.StatusBadRequest, "Bad Request", nil))
	}

	return ctx.Status(fiber.StatusCreated).
		JSON(model.APIResponse(nil, fiber.StatusCreated, "Created", response))
}

func (c *RegistrationController) UpdatePaymentStatus(ctx *fiber.Ctx) error {
	var request model.UpdatePaymentStatus

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(model.APIResponse("Cannot unmarshal body", fiber.StatusUnprocessableEntity, "Unprocessable Entity", nil))
	}

	errs := request.Validate()
	if errs != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.APIResponse(errs, fiber.StatusBadRequest, "Unprocessable Entity", nil))
	}

	status, err := c.RegistrationService.UpdatePaymentStatus(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).
			JSON(model.APIResponse(err.Error(), fiber.StatusBadRequest, "Bad Request", nil))
	}

	return ctx.Status(fiber.StatusOK).
		JSON(model.APIResponse("Updated", fiber.StatusOK, "Ok", map[string]string{
			"payment_status": status,
		}))
}
