package service

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/util"
	"github.com/google/uuid"
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
	// check email if already exists
	respEmail, _ := s.RegistrationRepository.GetByEmail(request.Email)
	if respEmail != nil {
		return nil, errors.New("email has been recorded")
	}

	// check phone number if already exists
	respPhone, _ := s.RegistrationRepository.GetByPhone(request.Phone)
	if respPhone != nil {
		return nil, errors.New("phone has been recorded")
	}

	username := uuid.NewString()
	password := uuid.NewString()
	passwordHash, _ := util.Hash(password)

	var register *entity.Registration

	if program == entity.S1D3D4 {
		register = entity.NewRegisterS1D3D4(
			uuid.NewString(),
			request.Name,
			request.Email,
			request.Phone,
			username,
			passwordHash,
			time.Now().String(),
		)
	} else {
		register = entity.NewRegisterS2(
			uuid.NewString(),
			request.Name,
			request.Email,
			request.Phone,
			username,
			passwordHash,
			time.Now().String(),
		)
	}

	err := s.RegistrationRepository.Insert(register)
	if err != nil {
		log.Printf("Service.Create: %v", err.Error())
		return nil, err
	}

	response := model.RegistrationResponse{
		Username:      username,
		Password:      password,
		Bill:          register.Bill,
		AccountNumber: register.AccountNumber,
	}

	return &response, nil
}
