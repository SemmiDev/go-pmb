package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-clean/internal/app/model"
	"go-clean/internal/app/service"
	exception2 "go-clean/internal/exception"
)

type StudentController struct {
	StudentService service.StudentService
}

func NewStudentController(studentService *service.StudentService) StudentController {
	return StudentController{StudentService: *studentService}
}

func (controller *StudentController) Route(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Post("/login", controller.Login)

	api.Post("/students", controller.Create)
	api.Get("/students", controller.List)
	api.Get("/students/:id", controller.GetById)
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

func (controller *StudentController) Create(c *fiber.Ctx) error {
	var request model.CreateStudentRequest
	err := c.BodyParser(&request)
	request.ID = uuid.New().String()

	exception2.PanicIfNeeded(err)

	response := controller.StudentService.Create(&request)
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

func (controller *StudentController) Login(c *fiber.Ctx) error {
	var request model.AuthRequest
	err := c.BodyParser(&request)
	exception2.PanicIfNeeded(err)

	token := controller.StudentService.Login(c, &request)
	if token != nil {
		return c.JSON(model.WebResponse{
			Code:   fiber.StatusOK,
			Status: "ok",
			Data:   token,
		})
	}

	return c.JSON(model.WebResponse{
		Code:   fiber.StatusUnauthorized,
		Status: "unauthorozed",
		Data:   "unauthorozed",
	})
}
