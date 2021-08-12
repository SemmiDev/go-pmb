package interfaces

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/requests"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/responses"
)

type IRegistrationService interface {
	Register(request *requests.Register) *responses.HttpResponse
	UpdatePaymentStatus(request *requests.UpdatePaymentStatus) *responses.HttpResponse
}
