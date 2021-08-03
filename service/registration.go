package service

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/constant"
	"github.com/SemmiDev/fiber-go-clean-arch/domain"
	"github.com/SemmiDev/fiber-go-clean-arch/mailer"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/util"
	"github.com/twinj/uuid"
	"log"
	"time"
)

type service struct {
	RegistrationRepository domain.RegistrationRepository
	Mailer                 mailer.Mailer
}

func NewRegistrationService(registrationRepo *domain.RegistrationRepository, mailService *mailer.Mailer) domain.RegistrationService {
	return &service{
		RegistrationRepository: *registrationRepo,
		Mailer:                 *mailService,
	}
}

func (s *service) Create(request *model.RegistrationRequest, program constant.Program) (*model.RegistrationResponse, error) {
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
	passwordHash, err := util.Hash(password)
	if err != nil {
		log.Printf("Service.Hash: %v", err.Error())
		return nil, err
	}
	va := util.RandomVirtualAccount(request.Phone)

	// define bill by program
	var bill = constant.S1D3D4Bill
	if program == constant.S2 {
		bill = constant.S2Bill
	}

	// payload
	var register = domain.NewRegistration(
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

	response := model.RegistrationResponse{
		Recipient:      register.Email,
		Username:       register.Username,
		Password:       password,
		VirtualAccount: register.VirtualAccount,
		Bill:           register.Bill,
	}

	s.Mailer.SendEmail(constant.RegistrationTemplate, &response)
	if err != nil {
		log.Printf("Service.SendEmail: %v", err.Error())
		return nil, err
	}
	return &response, nil
}

func (s *service) GetByUsername(req *model.LoginRequest) (*domain.Registration, error) {
	exists, err := s.RegistrationRepository.GetByUsername(req)
	if err != nil {
		return nil, err
	}
	return exists, nil
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
