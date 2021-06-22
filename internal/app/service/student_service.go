package service

import (
	"github.com/gofiber/fiber/v2"
	"go-clean/internal/app/model"
)

type StudentService interface {
	Login(c *fiber.Ctx, request *model.AuthRequest) *model.Token
	Create(request *model.CreateStudentRequest) (response *model.CreateStudentResponse)
	Get(id string)  (response *model.GetSingleStudentResponse)
	List() (responses []model.GetStudentResponse)
}