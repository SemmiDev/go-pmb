package command

import (
	"github.com/SemmiDev/go-pmb/internal/registrant/domain"
)

type RegistrantCommand interface {
	Save(registrant *domain.Registrant) <-chan error
	UpdateStatus(id string, status domain.PaymentStatus) <-chan error
}
