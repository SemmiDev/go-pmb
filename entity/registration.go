package entity

import (
	"github.com/SemmiDev/fiber-go-clean-arch/model"
)

type Registration struct {
	ID            string        `bson:"id"`
	Name          string        `bson:"name"`
	Email         string        `bson:"email"`
	Phone         string        `bson:"phone"`
	Username      string        `bson:"username"`
	Password      string        `bson:"password"`
	Kind          Program       `bson:"kind"`
	Bill          Bill          `bson:"bill"`
	AccountNumber AccountNumber `bson:"account_number"`
	Status        bool          `bson:"status"`
	CreatedAt     string        `bson:"created_at"`
}

func NewRegisterS1D3D4(id, name, email, phone, username, passwordHash, time string) *Registration {
	return &Registration{
		ID:            id,
		Name:          name,
		Email:         email,
		Phone:         phone,
		Username:      username,
		Password:      passwordHash,
		Kind:          S1D3D4,
		Bill:          S1D3D4Bill,
		AccountNumber: S1D3D4AccountNumber,
		Status:        false,
		CreatedAt:     time,
	}
}

func NewRegisterS2(id, name, email, phone, username, passwordHash, time string) *Registration {
	return &Registration{
		ID:            id,
		Name:          name,
		Email:         email,
		Phone:         phone,
		Username:      username,
		Password:      passwordHash,
		Kind:          S2,
		Bill:          S2Bill,
		AccountNumber: S2AccountNumber,
		Status:        false,
		CreatedAt:     time,
	}
}

type RegistrationRepository interface {
	Insert(register *Registration) error
	GetByEmail(email string) (*Registration, error)
	GetByPhone(phone string) (*Registration, error)
	DeleteAll()
}

type RegistrationService interface {
	Create(request *model.RegistrationRequest, program Program) (*model.RegistrationResponse, error)
}
