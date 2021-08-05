package service

import (
	"encoding/json"
	"github.com/SemmiDev/fiber-go-clean-arch/constant"
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/util"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/twinj/uuid"
	"go.uber.org/zap"
	"sync"
	"time"
)

type service struct {
	RegistrationRepository entity.RegistrationRepository
	MailBroker             *amqp.Channel
}

func NewRegistrationService(registrationRepo *entity.RegistrationRepository, mailBroker *amqp.Channel) entity.RegistrationService {
	return &service{
		RegistrationRepository: *registrationRepo,
		MailBroker:             mailBroker,
	}
}

var Error error

func (s *service) Create(request *model.RegistrationRequest, program constant.Program) (*model.RegistrationResponse, error) {
	// make sure when a new request is coming, set Errors to nil
	Error = nil

	var wg sync.WaitGroup
	wg.Add(2)
	go s.RegistrationRepository.GetByEmail(&wg, request.Email)
	go s.RegistrationRepository.GetByPhone(&wg, request.Phone)
	wg.Wait()

	if Error != nil {
		return nil, Error
	}

	// prepare username, password, and generate va
	username, password := uuid.NewV4().String(), uuid.NewV4().String()
	passwordHash, err := util.Hash(password)
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	va := util.RandomVirtualAccount(request.Phone)

	// define bill by program
	var bill = constant.S1D3D4Bill
	if program == constant.S2 {
		bill = constant.S2Bill
	}

	// payload
	var register = entity.NewRegistration(
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
		zap.S().Error(err.Error())
		return nil, err
	}

	response := model.RegistrationResponse{
		Recipient:      register.Email,
		Username:       register.Username,
		Password:       password,
		VirtualAccount: register.VirtualAccount,
		Bill:           register.Bill,
	}

	payload, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}

	msg := amqp.Publishing{
		ContentType: fiber.MIMEApplicationJSON,
		Body:        payload,
	}

	if err = s.MailBroker.Publish(
		"",
		"QueueEmailServiceRegistration",
		false,
		false,
		msg,
	); err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}

	return &response, nil
}

func (s *service) GetByUsername(req *model.LoginRequest) (*entity.Registration, error) {
	exists, err := s.RegistrationRepository.GetByUsername(req)
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	return exists, nil
}

func (s *service) UpdateStatusBilling(va *model.UpdateStatus) error {
	exists, err := s.RegistrationRepository.GetByVa(va)
	if err != nil {
		zap.S().Error(err.Error())
		return err
	}
	err = s.RegistrationRepository.UpdateStatus(exists.VirtualAccount)
	if err != nil {
		zap.S().Error(err.Error())
		return err
	}
	return nil
}
