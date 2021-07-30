package service

import (
	"errors"
	"github.com/google/uuid"
	"go-clean/entity"
	"go-clean/model"
	"go-clean/util"
	"log"
	"time"
)

type service struct {
	RegistrationRepository entity.RegistrationRepository
}

func NewRegistrationService(rp *entity.RegistrationRepository) entity.RegistrationService {
	return &service{
		RegistrationRepository: *rp,
	}
}

func (s *service) Create(request *model.RegistrationRequest, program entity.Program) (*model.RegistrationResponse, error) {
	// email number is exists
	respEmail, _ := s.RegistrationRepository.GetByEmail(request.Email)
	if respEmail != nil {
		return nil, errors.New("email has been recorded")
	}

	// check phone number is exists
	respPhone, _ := s.RegistrationRepository.GetByPhone(request.Phone)
	if respPhone != nil {
		return nil, errors.New("phone has been recorded")
	}

	username := uuid.NewString()
	password := uuid.NewString()
	passwordHash, _ := util.Hash(password)

	register := entity.Registration{
		ID:        uuid.NewString(),
		Name:      request.Name,
		Email:     request.Email,
		Phone:     request.Phone,
		Username:  username,
		Password:  passwordHash,
		Kind:      program,
		CreatedAt: time.Now().String(),
	}

	err := s.RegistrationRepository.Insert(&register)
	if err != nil {
		log.Printf("Service.Create: %v", err.Error())
		return nil, err
	}

	response := model.RegistrationResponse{
		Username: username,
		Password: password,
	}

	return &response, nil
}
