package repository

import (
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
)

type RegistrationRepository interface {
	Insert(register *entity.Registration) error
	GetByID(ID string) *entity.Registration
	GetByUsername(username string) *entity.Registration
	GetByEmail(email string) bool
	GetByPhone(phone string) bool
	UpdateStatus(ID string, status string) error
	DeleteAll() error
}
