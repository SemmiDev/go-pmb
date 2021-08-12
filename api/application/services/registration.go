package services

import (
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/requests"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/responses"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/mappings"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/constants"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/customErrors"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/adapters"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/payments/midtrans"
	"github.com/SemmiDev/fiber-go-clean-arch/notifier/environments"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegistrationService struct {
	HashAdapter            adapters.IHashAdapter
	Midtrans               midtrans.IMidtrans
	RabbitMQ               interfaces.IRabbitMQ
	RegistrationRepository interfaces.IRegistrationRepository
}

func NewRegistrationService(
	h adapters.IHashAdapter,
	m midtrans.IMidtrans,
	q interfaces.IRabbitMQ,
	r interfaces.IRegistrationRepository,

) interfaces.IRegistrationService {
	return &RegistrationService{
		HashAdapter:            h,
		Midtrans:               m,
		RabbitMQ:               q,
		RegistrationRepository: r,
	}
}

func (r *RegistrationService) Register(request *requests.Register) *responses.HttpResponse {
	// Check email is already exists or not
	registerByEmail, _ := r.RegistrationRepository.GetByEmail(request.Email)
	if registerByEmail != nil {
		return responses.BadRequest(customErrors.EmailAlreadyExistsMessage)
	}

	// Phone number is already exists or not
	registerByPhone, _ := r.RegistrationRepository.GetByEmail(request.Email)
	if registerByPhone != nil {
		return responses.BadRequest(customErrors.PhoneAlreadyExistsMessage)
	}

	// Generate username & password
	username := adapters.NewRandomAdapter().GenerateRandom()
	password := adapters.NewRandomAdapter().GenerateRandom()

	// Hash password
	passwordHash, err := r.HashAdapter.GenerateHash(password)
	if err != nil {
		return responses.BadRequest(err.Error())
	}

	register := mappings.NewRegistrationTemplate(
		request.Name,
		request.Email,
		username,
		passwordHash,
		request.Phone,
		request.Program)

	id := adapters.NewUuidAdapter().GenerateUuid()
	code := fmt.Sprintf("%s-%s", "register", id)
	register.Code = code

	// Define payments
	payment := midtrans.Payment{
		ID:     id,
		Amount: register.Bill,
	}

	// Generate payments URL
	paymentURL, err := r.Midtrans.GetPaymentURL(&payment, register)
	if err != nil {
		return responses.InternalServerError(err)
	}
	register.PaymentURL = paymentURL

	result, err := r.RegistrationRepository.Insert(register)
	if err != nil {
		return responses.InternalServerError(err)
	}

	response := responses.RegisterResponse{
		Recipient:  result.Email,
		Username:   username,
		Password:   password,
		Bill:       payment.AmountFormatIDR(),
		PaymentURL: paymentURL,
	}

	err = r.RabbitMQ.SendMessage(environments.RabbitMQQueue, &response)
	if err != nil {
		return responses.InternalServerError(err)
	}

	return responses.Created(response)
}

func (r *RegistrationService) UpdatePaymentStatus(request *requests.UpdatePaymentStatus) *responses.HttpResponse {
	// Check register id is already exists or not
	ID, _ := primitive.ObjectIDFromHex(request.RegisterID)
	register, err := r.RegistrationRepository.GetByID(ID)
	if register == nil {
		return responses.InternalServerError(err)
	}

	if request.PaymentType == "credit_card" && request.PaymentStatus == "capture" && request.FraudStatus == "accept" {
		register.Status = constants.PaymentStatusPaid
	} else if request.PaymentStatus == "settlement" {
		register.Status = constants.PaymentStatusPaid
	} else if request.PaymentStatus == "deny" || request.PaymentStatus == "expire" || request.PaymentStatus == "cancel" {
		register.Status = constants.PaymentStatusCancel
	}

	err = r.RegistrationRepository.UpdateStatus(register.ID, register.Status)
	if err != nil {
		return responses.InternalServerError(err)
	}

	return responses.Ok("Payment Status Updated")
}
