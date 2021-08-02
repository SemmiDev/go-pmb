package service

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/mailer/tasks"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/util"
	"github.com/hibiken/asynq"
	"github.com/twinj/uuid"
	"log"
	"time"
)

type service struct {
	RegistrationRepository model.RegistrationRepository
	EmailAsynq             *asynq.Client
}

func NewRegistrationService(
	rp *model.RegistrationRepository,
	EmailAsynq *asynq.Client) model.RegistrationService {

	return &service{
		RegistrationRepository: *rp,
		EmailAsynq:             EmailAsynq,
	}
}

func (s *service) Create(request *model.RegistrationRequest, program model.Program) (*model.RegistrationResponse, error) {
	// check mailer if already exists
	respEmail, _ := s.RegistrationRepository.GetByEmail(request.Email)
	if respEmail != nil {
		return nil, errors.New("mailer has been recorded")
	}

	// check phone number if already exists
	respPhone, _ := s.RegistrationRepository.GetByPhone(request.Phone)
	if respPhone != nil {
		return nil, errors.New("phone has been recorded")
	}

	username := uuid.NewV4().String()
	password := uuid.NewV4().String()
	passwordHash, _ := util.Hash(password)
	va := util.RandomVirtualAccount(request.Phone)

	var register *model.Registration
	if program == model.S1D3D4 {
		register = model.RegisterS2PrototypePrototype()
		register.ID = uuid.NewV4().String()
		register.Name = request.Name
		register.Email = request.Email
		register.Phone = request.Phone
		register.Username = username
		register.Password = passwordHash
		register.VirtualAccount = va
		register.CreatedAt = time.Now().String()
	} else {
		register = model.RegisterS2PrototypePrototype()
		register.ID = uuid.NewV4().String()
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

	// sent mail
	task, err := tasks.NewRegisterEmail(
		response.Username,
		response.Password,
		register.Email,
		response.Bill,
		response.VirtualAccount,
	)
	if err != nil {
		return nil, err
	}

	// Process the task immediately in critical queue.
	_, err = s.EmailAsynq.Enqueue(task)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *service) GetByUsername(req *model.LoginRequest) (*model.Registration, error) {
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
