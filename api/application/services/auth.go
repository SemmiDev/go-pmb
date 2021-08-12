package services

import (
	"context"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/requests"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/responses"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/constants"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/customErrors"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/adapters"
)

type AuthService struct {
	RegisterRepository interfaces.IRegistrationRepository
	AuthRepository     interfaces.IRedisAuthRepository
	HashAdapter        adapters.IHashAdapter
	JwtAdapter         adapters.IJwtAdapter
}

func NewAuthService(
	r interfaces.IRegistrationRepository,
	a interfaces.IRedisAuthRepository,
	h adapters.IHashAdapter,
	j adapters.IJwtAdapter) interfaces.IAuthService {

	return &AuthService{RegisterRepository: r, AuthRepository: a, HashAdapter: h, JwtAdapter: j}
}

func (a *AuthService) Login(request *requests.Login) *responses.HttpResponse {
	register, err := a.RegisterRepository.GetByUsername(request.Username)
	if err != nil || register == nil {
		return responses.InternalServerError(err)
	}

	if isValid := a.HashAdapter.CheckHash(register.Password, request.Password); !isValid {
		return responses.BadRequest(customErrors.InvalidPasswordMessage)
	}

	if register.Status != constants.PaymentStatusPaid {
		return responses.BadRequest(customErrors.NotYetBIllMessage)
	}

	token, err := a.JwtAdapter.CreateTokenJWT(register.ID.Hex(), register.Email)
	if err != nil {
		return responses.InternalServerError(err)
	}

	saveErr := a.AuthRepository.CreateAuth(context.Background(), register.ID.Hex(), token)
	if saveErr != nil {
		return responses.InternalServerError(err)
	}

	response := responses.NewLoginResponse(
		register.ID.Hex(),
		register.Name,
		register.Username,
		token.AccessToken,
		token.RefreshToken,
	)
	return responses.Ok(response)
}
