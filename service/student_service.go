package service

import (
	"go-clean/model"
)

type StudentService interface {
	Create(request model.CreateStudentRequest) (response model.CreateStudentResponse)
	List() (responses []model.GetStudentResponse)
	Delete(id string) string
	Get(id string)  (response model.GetSingleStudentResponse)
	Update(id string, student model.UpdateStudentRequest) bool
}