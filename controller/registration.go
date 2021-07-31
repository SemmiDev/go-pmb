package controller

import (
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/gofiber/fiber/v2"
)

type RegistrationController struct {
	RegistrationService entity.RegistrationService
}

func NewRegistrationController(registrationService *entity.RegistrationService) RegistrationController {
	return RegistrationController{
		RegistrationService: *registrationService,
	}
}

func (c *RegistrationController) Route(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/registration", c.Create)
}

func (c *RegistrationController) Create(ctx *fiber.Ctx) error {
	var request model.RegistrationRequest

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(model.WebResponse{
			Code:         fiber.StatusUnprocessableEntity,
			Status:       "Unprocessable Entity",
			Error:        true,
			ErrorMessage: "Cannot unmarshal body",
			Data:         nil,
		})
	}

	errs := request.Validate()
	if errs != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: errs,
			Data:         nil,
		})
	}

	var program entity.Program
	switch request.Program {
	case "S1D3D4":
		program = entity.S1D3D4
	case "S2":
		program = entity.S2
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Bad Request",
			Error:  true,
			ErrorMessage: map[string]string{
				"Program_Not_Available": "Please Chose Between S1D3D4 or S2",
			},
			Data: nil,
		})
	}

	response, err := c.RegistrationService.Create(&request, program)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse{
			Code:         fiber.StatusBadRequest,
			Status:       "Bad Request",
			Error:        true,
			ErrorMessage: err.Error(),
			Data:         nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.WebResponse{
		Code:         fiber.StatusCreated,
		Status:       "Created",
		Error:        false,
		ErrorMessage: nil,
		Data:         response,
	})
}
