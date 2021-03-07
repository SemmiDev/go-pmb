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

