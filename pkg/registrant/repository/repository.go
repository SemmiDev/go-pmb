package repository

import (
	"github.com/SemmiDev/go-pmb/pkg/registrant/domain"
)

type RegistrantRepository interface {
	Save(registrant *domain.Registrant) <-chan error
	UpdateStatus(id string, status domain.PaymentStatus) <-chan error
}
