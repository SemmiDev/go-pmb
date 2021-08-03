package fakeservice

import (
	"errors"
	model2 "github.com/SemmiDev/fiber-go-clean-arch/internal/model"
	util2 "github.com/SemmiDev/fiber-go-clean-arch/pkg/util"
	"github.com/twinj/uuid"
	"log"
	"time"
)

// without email support for controller test

type service struct {
	RegistrationRepository model2.RegistrationRepository
}

func NewRegistrationService(registrationRepo *model2.RegistrationRepository) model2.RegistrationService {
	return &service{
		RegistrationRepository: *registrationRepo,
	}
}

func (s *service) Create(request *model2.RegistrationRequest, program model2.Program) (*model2.RegistrationResponse, error) {
	// check mailer if already exists
	respEmail, _ := s.RegistrationRepository.GetByEmail(request.Email)
	if respEmail != nil {
		return nil, errors.New("email has been recorded")
	}

	// check phone number if already exists
	respPhone, _ := s.RegistrationRepository.GetByPhone(request.Phone)
	if respPhone != nil {
		return nil, errors.New("phone has been recorded")
	}

	// prepare username, password, and generate va
	username, password := uuid.NewV4().String(), uuid.NewV4().String()
	passwordHash, err := util2.Hash(password)
	if err != nil {
		log.Printf("Service.Hash: %v", err.Error())
		return nil, err
	}
	va := util2.RandomVirtualAccount(request.Phone)

	// define bill by program
	var bill = model2.S1D3D4Bill
	if program == model2.S2 {
		bill = model2.S2Bill
	}

	// payload
	var register = model2.NewRegistration(
		uuid.NewV4().String(),
		request.Name,
		request.Email,
		request.Phone,
		username,
		passwordHash,
		program,
		bill,
		va,
		false,
		time.Now().String(),
	)

	err = s.RegistrationRepository.Insert(register)
	if err != nil {
		log.Printf("Service.Create: %v", err.Error())
		return nil, err
	}

	response := model2.RegistrationResponse{
		Recipient:      register.Email,
		Username:       register.Username,
		Password:       password,
		VirtualAccount: register.VirtualAccount,
		Bill:           register.Bill,
	}

	return &response, nil
}

func (s *service) GetByUsername(req *model2.LoginRequest) (*model2.Registration, error) {
	exists, err := s.RegistrationRepository.GetByUsername(req)
	if err != nil {
		return nil, err
	}
	return exists, nil
}

func (s *service) UpdateStatusBilling(va *model2.UpdateStatus) error {
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
