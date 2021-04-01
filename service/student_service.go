package service

import "go-clean/model"

type StudentService interface {
	Create(request model.CreateStudentRequest) (response model.CreateStudentResponse)
	List() (responses []model.GetStudentResponse)
	Delete(id string) string
}