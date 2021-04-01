package service

import (
	"go-clean/entity"
	"go-clean/model"
	"go-clean/repository"
	validation "go-clean/validation"
)

func NewStudentService(studentRepository *repository.StudentRepository) StudentService {
	return &studentServiceImpl{
		StudentRepository: *studentRepository,
	}
}

type studentServiceImpl struct {
	StudentRepository repository.StudentRepository
}

func (service *studentServiceImpl) Update(id string, student model.UpdateStudentRequest) bool {
	std := entity.Student{
		Id:         "",
		Identifier: "",
		Name:       student.Name,
		Email:      student.Email,
	}
	return service.StudentRepository.UpdateById(id, std)
}

func (service *studentServiceImpl) Get(id string) (response model.GetSingleStudentResponse) {
	student := service.StudentRepository.GetById(id)
	return model.GetSingleStudentResponse{
		Identifier: student.Identifier,
		Name:       student.Name,
		Email:      student.Email,
	}
}

func (service *studentServiceImpl) Delete(id string) string {
	return service.StudentRepository.Delete(id)
}

func (service *studentServiceImpl) Create(request model.CreateStudentRequest) (response model.CreateStudentResponse) {
	validation.Validate(request)

	student := entity.Student{
		Id:         request.Id,
		Identifier: request.Identifier,
		Name:       request.Name,
		Email:      request.Email,
	}

	service.StudentRepository.Insert(student)

	response = model.CreateStudentResponse{
		Id:         student.Id,
		Name:       student.Name,
		Identifier: student.Identifier,
		Email:      student.Email,
	}
	return response
}

func (service *studentServiceImpl) List() (responses []model.GetStudentResponse) {
	students := service.StudentRepository.FindAll()
	for _, student := range students {
		responses = append(responses, model.GetStudentResponse{
			Id:         student.Id,
			Name:       student.Name,
			Identifier: student.Identifier,
			Email:      student.Email,
		})
	}
	return responses
}
