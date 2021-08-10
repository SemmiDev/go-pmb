package services

import (
	"encoding/json"
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/constant"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/entities"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/helper"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/models"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/twinj/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type RegistrationService interface {
	Register(m *models.RegistrationRequest) (*models.RegistrationResponse, error)
	UpdatePaymentStatus(m *models.UpdatePaymentStatusRequest) (string, error)
	Login(m *models.LoginRequest) (*entities.Registration, error)
}

type registrationService struct {
	RegistrationRepository repositories.RegistrationRepository
	MailBroker             *amqp.Channel
	PaymentService         Service
}

func NewRegistrationService(rp *repositories.RegistrationRepository, mb *amqp.Channel, ps Service) RegistrationService {
	return &registrationService{
		RegistrationRepository: *rp,
		MailBroker:             mb,
		PaymentService:         ps,
	}
}

func (r *registrationService) Register(register *models.RegistrationRequest) (*models.RegistrationResponse, error) {
	// Check email is already exists or not
	emailExists := r.RegistrationRepository.GetByEmail(register.Email)
	if emailExists {
		return nil, constant.ErrEmail
	}

	// Check phone is already exists or not
	phoneExists := r.RegistrationRepository.GetByPhone(register.Phone)
	if phoneExists {
		return nil, constant.ErrPhone
	}

	// Generate username & password
	username, password := helper.Random(), helper.Random()

	// Hash password
	passwordHash, err := helper.Hash(password)
	if err != nil {
		return nil, err
	}

	// Define a registerData (default s1)
	var registerData = entities.RegisterFactory(register.Program)

	// Assign other data
	registerData.ID = uuid.NewV4().String()
	registerData.Name = register.Name
	registerData.Username = username
	registerData.Password = passwordHash
	registerData.Email = register.Email
	registerData.Phone = register.Phone

	// Time
	now := primitive.NewDateTimeFromTime(time.Now())
	registerData.CreatedAt = now
	registerData.UpdatedAt = now

	codeValue := "REGISTER"
	code := fmt.Sprintf("%s-%s", codeValue, registerData.ID)
	registerData.Code = code

	// Define payment
	payment := models.Payment{
		ID:     registerData.ID,
		Amount: registerData.Bill,
	}

	// Generate payment URL
	paymentURL, err := r.PaymentService.GetPaymentURL(&payment, registerData)
	if err != nil {
		return nil, err
	}
	registerData.PaymentURL = paymentURL

	// Store to db
	err = r.RegistrationRepository.Insert(registerData)
	if err != nil {
		return nil, err
	}

	// Create the response
	response := models.NewRegistrationResponse(
		registerData.Email,
		registerData.Username,
		password,
		payment.AmountFormatIDR(),
		registerData.PaymentURL,
	)

	// Uncomment if need sent to email also (nb: this section using msg broker, so you need to enable mail service first)
	//err = r.SendEmail(response)
	//if err != nil {
	//	return nil, err
	//}

	return response, nil
}

func (r *registrationService) UpdatePaymentStatus(input *models.UpdatePaymentStatusRequest) (string, error) {
	register := r.RegistrationRepository.GetByID(input.RegisterID)
	if register == nil {
		return "", constant.ErrIDNotFound
	}

	if input.PaymentType == "credit_card" && input.PaymentStatus == "capture" && input.FraudStatus == "accept" {
		register.Status = constant.PaymentStatusPaid
	} else if input.PaymentStatus == "settlement" {
		register.Status = constant.PaymentStatusPaid
	} else if input.PaymentStatus == "deny" || input.PaymentStatus == "expire" || input.PaymentStatus == "cancel" {
		register.Status = constant.PaymentStatusCancel
	}

	err := r.RegistrationRepository.UpdateStatus(register.ID, register.Status)
	if err != nil {
		return "", err
	}

	return register.Status, nil
}

func (r *registrationService) SendEmail(response models.RegistrationResponse) error {
	// Marshal the response
	payload, err := json.Marshal(response)
	if err != nil {
		return err
	}

	// Msg payload to mail broker
	msg := amqp.Publishing{
		ContentType: fiber.MIMEApplicationJSON,
		Body:        payload,
	}

	// Publish to mail broker
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

func (r *registrationService) Login(m *models.LoginRequest) (*entities.Registration, error) {
	register := r.RegistrationRepository.GetByUsername(m.Username)
	if register == nil {
		if register == nil {
			return nil, constant.ErrRegisterNotFound
		}
	}

	err := helper.Check(m.Password, register.Password)
	if err != nil {
		return nil, err
	}

	if register.Status != constant.PaymentStatusPaid {
		return nil, constant.ErrBillHasNotBeenPaid
	}

	return register, nil
}
