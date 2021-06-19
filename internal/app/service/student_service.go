package service

import (
	"go-clean/internal/app/model"
)

type StudentService interface {
	Create(request *model.CreateStudentRequest) (response *model.CreateStudentResponse)
	Get(id string)  (response *model.GetSingleStudentResponse)
	List() (responses []model.GetStudentResponse)
}