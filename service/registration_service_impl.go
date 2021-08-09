package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/constant"
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/SemmiDev/fiber-go-clean-arch/payment"
	"github.com/SemmiDev/fiber-go-clean-arch/repository"
	"github.com/SemmiDev/fiber-go-clean-arch/util"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type registrationService struct {
	RegistrationRepository repository.RegistrationRepository
	MailBroker             *amqp.Channel
	PaymentService         payment.Service
}

func NewRegistrationService(rp *repository.RegistrationRepository, mb *amqp.Channel, ps payment.Service) RegistrationService {
	return &registrationService{
		RegistrationRepository: *rp,
		MailBroker:             mb,
		PaymentService:         ps,
	}
}

var (
	ErrEmail    = errors.New("e-mail has been registered")
	ErrPhone    = errors.New("phone has been registered")
	ErrNotFound = errors.New("ID not found")
)

func (r *registrationService) Register(register *model.RegistrationRequest) (*model.RegistrationResponse, error) {
	// check email is already exists or not
	emailExists := r.RegistrationRepository.GetByEmail(register.Email)
	if emailExists {
		return nil, constant.ErrEmail
	}
	// check phone is already exists or not
	phoneExists := r.RegistrationRepository.GetByPhone(register.Phone)
	if phoneExists {
		return nil, constant.ErrPhone
	}

	// generate username & password
	username, password := util.Random(), util.Random()
	// hash password
	passwordHash, err := util.Hash(password)
	if err != nil {
		return nil, err
	}

	// define a registerData (default s1)
	var registerData = entity.RegisterS1D3D4Prototype
	if register.Program == constant.S2 {
		registerData = entity.RegisterS2Prototype
	}

	// assign other data
	registerData.ID = uuid.NewV4().String()
	registerData.Name = register.Name
	registerData.Username = username
	registerData.Password = passwordHash
	registerData.Email = register.Email
	registerData.Phone = register.Phone

	// time
	now := primitive.NewDateTimeFromTime(time.Now())
	registerData.CreatedAt = now
	registerData.UpdatedAt = now

	codeValue := "REGISTRATION"
	code := fmt.Sprintf("%s-%s", codeValue, registerData.ID)
	registerData.Code = code

	// define payment
	payment := payment.Payment{
		ID:     registerData.ID,
		Amount: registerData.Bill,
	}

	// generate payment URL
	paymentURL, err := r.PaymentService.GetPaymentURL(&payment, registerData)
	if err != nil {
		return nil, err
	}
	registerData.PaymentURL = paymentURL

	// store to db
	err = r.RegistrationRepository.Insert(registerData)
	if err != nil {
		return nil, err
	}

	// create the response
	response := model.RegistrationResponse{
		Recipient:  registerData.Email,
		Username:   registerData.Username,
		Password:   password,
		Bill:       payment.AmountFormatIDR(),
		PaymentURL: registerData.PaymentURL,
	}

	// uncomment if need sent to email also
	//err = r.SendEmail(response)
	//if err != nil {
	//	return nil, err
	//}

	return &response, nil
}

func (r *registrationService) UpdatePaymentStatus(input *model.UpdatePaymentStatus) (string, error) {
	register := r.RegistrationRepository.GetByID(input.RegisterID)
	if register == nil {
		return "", constant.ErrIDNotFound
	}

	if input.PaymentType == "credit_card" && input.PaymentStatus == "capture" && input.FraudStatus == "accept" {
		register.Status = constant.PAID
	} else if input.PaymentStatus == "settlement" {
		register.Status = constant.PAID
	} else if input.PaymentStatus == "deny" || input.PaymentStatus == "expire" || input.PaymentStatus == "cancel" {
		register.Status = constant.CANCEL
	}

	err := r.RegistrationRepository.UpdateStatus(register.ID, register.Status)
	if err != nil {
		return "", err
	}

	return register.Status, nil
}

func (r *registrationService) SendEmail(response model.RegistrationResponse) error {
	// marshal the response
	payload, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// msg payload to mail broker
	msg := amqp.Publishing{
		ContentType: fiber.MIMEApplicationJSON,
		Body:        payload,
	}

	// publish to mail broker
	if err = r.MailBroker.Publish(
		"",
		"QueueEmailServiceRegistration",
		false,
		false,
		msg,
	); err != nil {
		return err
	}

	return nil
}

func (r *registrationService) Login(m *model.LoginRequest) (*entity.Registration, error) {
	register := r.RegistrationRepository.GetByUsername(m.Username)
	if register == nil {
		if register == nil {
			return nil, constant.ErrRegisterNotFound
		}
	}

	err := util.Check(m.Password, register.Password)
	if err != nil {
		return nil, err
	}

	if register.Status != constant.PAID {
		return nil, constant.ErrBillHasNotBeenPaid
	}

	return register, nil
}
