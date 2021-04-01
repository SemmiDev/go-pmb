package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-clean/exception"
	"go-clean/model"
	"go-clean/service"
)

type StudentController struct {
	StudentService service.StudentService
}

func NewStudentController(studentService *service.StudentService) StudentController {
	return StudentController{StudentService: *studentService}
}

func (controller *StudentController) Route(app *fiber.App) {
	app.Post("/api/students", controller.Create)
	app.Get("/api/students", controller.List)
	app.Delete("/api/students/:id", controller.Delete)
	app.Get("/api/students/:id", controller.GetById)
	app.Put("/api/students/:id", controller.UpdateById)
}

func (controller *StudentController) UpdateById(c *fiber.Ctx) error {
	var id string = c.Params("id")
	var request model.UpdateStudentRequest
	err := c.BodyParser(&request)
	exception.PanicIfNeeded(err)

	response := controller.StudentService.Update(id, request)
	if response == true {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "UPDATED",
			Data:   response,
		})
	}

	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "FAILED TO UPDATED",
		Data:   nil,
	})
}

func (controller *StudentController) GetById(c *fiber.Ctx) error {
	var id string = c.Params("id")
	response := controller.StudentService.Get(id)

	return c.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "ID " + id + " found",
		Data:   response,
	})
}

func (controller *StudentController) Delete(c *fiber.Ctx) error {
	var id string = c.Params("id")
	response := controller.StudentService.Delete(id)
	if response == "DELETED" {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "DELETED",
			Data:   nil,
		})
	}
	return c.JSON(model.WebResponse{
		Code:   fiber.StatusExpectationFailed,
		Status: "ID NOT FOUND",
		Data:   nil,
	})
}

func (controller *StudentController) Create(c *fiber.Ctx) error {
	var request model.CreateStudentRequest
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.PanicIfNeeded(err)

	response := controller.StudentService.Create(request)
	return c.JSON(model.WebResponse{
		Code:   201,
		Status: "CREATED",
		Data:   response,
	})
}

func (controller *StudentController) List(c *fiber.Ctx) error {
	responses := controller.StudentService.List()
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

