package interfaces

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/requests"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/responses"
)

type IAuthService interface {
	Login(request *requests.Login) *responses.HttpResponse
}
