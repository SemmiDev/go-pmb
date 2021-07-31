package service

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/util"
	"github.com/google/uuid"
	"log"
	"time"
)

type service struct {
	RegistrationRepository model.RegistrationRepository
}

func NewRegistrationService(rp *model.RegistrationRepository) model.RegistrationService {
	return &service{
		RegistrationRepository: *rp,
	}
}

func (s *service) Create(request *model.RegistrationRequest, program model.Program) (*model.RegistrationResponse, error) {
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

	response := model.RegistrationResponse{
		Username: username,
		Password: password,
	}

	if program == model.S1D3D4 {
		model.RegisterS1D3D4Prototype.ID = uuid.NewString()
		model.RegisterS1D3D4Prototype.Name = request.Name
		model.RegisterS1D3D4Prototype.Email = request.Email
		model.RegisterS1D3D4Prototype.Phone = request.Phone
		model.RegisterS1D3D4Prototype.Username = username
		model.RegisterS1D3D4Prototype.Password = passwordHash
		model.RegisterS1D3D4Prototype.CreatedAt = time.Now().String()
		response.Bill = model.RegisterS1D3D4Prototype.Bill
		response.AccountNumber = model.RegisterS1D3D4Prototype.AccountNumber

		err := s.RegistrationRepository.Insert(model.RegisterS1D3D4Prototype)
		if err != nil {
			log.Printf("Service.Create: %v", err.Error())
			return nil, err
		}
	} else {
		model.RegisterS2Prototype.ID = uuid.NewString()
		model.RegisterS2Prototype.Name = request.Name
		model.RegisterS2Prototype.Email = request.Email
		model.RegisterS2Prototype.Phone = request.Phone
		model.RegisterS2Prototype.Username = username
		model.RegisterS2Prototype.Password = passwordHash
		model.RegisterS2Prototype.CreatedAt = time.Now().String()
		response.Bill = model.RegisterS2Prototype.Bill
		response.AccountNumber = model.RegisterS2Prototype.AccountNumber

		err := s.RegistrationRepository.Insert(model.RegisterS2Prototype)
		if err != nil {
			log.Printf("Service.Create: %v", err.Error())
			return nil, err
		}
	}

	return &response, nil
}
