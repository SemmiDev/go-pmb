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
	va := util.RandomVirtualAccount(request.Phone)

	var register *model.Registration
	if program == model.S1D3D4 {
		register = model.RegisterS2PrototypePrototype()

		register.ID = uuid.NewString()
		register.Name = request.Name
		register.Email = request.Email
		register.Phone = request.Phone
		register.Username = username
		register.Password = passwordHash
		register.VirtualAccount = va
		register.CreatedAt = time.Now().String()
	} else {
		register = model.RegisterS2PrototypePrototype()

		register.ID = uuid.NewString()
		register.Name = request.Name
		register.Email = request.Email
		register.Phone = request.Phone
		register.Username = username
		register.Password = passwordHash
		register.VirtualAccount = va
		register.CreatedAt = time.Now().String()
	}

	err := s.RegistrationRepository.Insert(register)
	if err != nil {
		log.Printf("Service.Create: %v", err.Error())
		return nil, err
	}

	response := model.RegistrationResponse{
		Username:       username,
		Password:       password,
		VirtualAccount: va,
		Bill:           register.Bill,
	}

	return &response, nil
}

func (s *service) UpdateStatusBilling(va *model.UpdateStatus) error {
	exists, err := s.RegistrationRepository.GetByVa(va)
	if err != nil {
		return err
	}
	err = s.RegistrationRepository.UpdateStatus(exists.VirtualAccount)
	if err != nil {
		log.Printf("Service.UpdateStatusBilling: %v \n", err.Error())
		return err
	}
	return nil
}
