package service

import (
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
)

type RegistrationService interface {
	Register(m *model.RegistrationRequest) (*model.RegistrationResponse, error)
	UpdatePaymentStatus(m *model.UpdatePaymentStatus) (string, error)
	Login(m *model.LoginRequest) (*entity.Registration, error)
}
