package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go-clean/internal/app/entity"
	"go-clean/internal/app/model"
	"go-clean/internal/app/repository"
	"go-clean/internal/auth"
	"go-clean/internal/exception"
	"go-clean/internal/util"
	"go-clean/internal/validation"
	"strings"
)

func NewStudentService(studentRepository *repository.StudentRepository) StudentService {
	return &service{
		StudentRepository: *studentRepository,
	}
}

type service struct {
	StudentRepository repository.StudentRepository
}

func (s *service) Login(c *fiber.Ctx, request *model.AuthRequest) *model.Token {
	validation.ValidateLogin(*request)
	student, err := s.StudentRepository.GetByIdentifier(request.Identifier)
	exception.PanicIfNeeded(err)

	if request.Password == student.Password {
		ts, err := auth.CreateToken(student.Id)
		exception.PanicIfNeeded(err)

		saveErr := auth.CreateAuth(c, student.Id, ts)
		exception.PanicIfNeeded(saveErr)

		return &model.Token{
			AccessToken: ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
	}
	return nil
}

func (s *service) Create(request *model.CreateStudentRequest) (response *model.CreateStudentResponse) {
	validation.Validate(*request)

	// path
	var path string
	switch request.Path {
	case 1:
		path = "SNMPTN"
	case 2:
		path = "SBMPTN"
	case 3:
		path = "PBUD"
	}

	identifierGen, emailStudentGen := util.NIMAndEmailGenerator(request.FullName, request.Path, request.Year)
	student := entity.Student{
		Id: request.ID,
		Identifier:         identifierGen,
		FullName:           strings.Title(request.FullName),
		Email:              strings.ToLower(request.Email),
		EmailStudent:       emailStudentGen,
		PhoneNumber:        request.PhoneNumber,
		Path:               path,
		RegistrationNumber: request.RegistrationNumber,
		Password:           uuid.NewString(),
	}

	oid := s.StudentRepository.Insert(&student)
	result, err := s.StudentRepository.GetByOID(oid)
	exception.PanicIfNeeded(err)

	response = &model.CreateStudentResponse{
		Identifier: result.Identifier,
		Password:   result.Password,
	}
	return
}

func (s *service) Get(id string) (response *model.GetSingleStudentResponse) {
	student := s.StudentRepository.GetById(id)
	return &model.GetSingleStudentResponse{
		Id:                 student.Id,
		Identifier:         student.Identifier,
		FullName:           student.FullName,
		Email:              student.Email,
		EmailStudent:       student.EmailStudent,
		PhoneNumber:        student.PhoneNumber,
		Path:               student.Path,
		RegistrationNumber: student.RegistrationNumber,
		Password:           student.Password,
	}
}

func (s *service) List() (responses []model.GetStudentResponse) {
	students := s.StudentRepository.FindAll()
	for _, student := range students {
		responses = append(responses, model.GetStudentResponse{
			Id:                 student.Id,
			Identifier:         student.Identifier,
			FullName:           student.FullName,
			Email:              student.Email,
			EmailStudent:       student.EmailStudent,
			PhoneNumber:        student.PhoneNumber,
			Path:               student.Path,
			RegistrationNumber: student.RegistrationNumber,
		})
	}
	return responses
}