package entity

import (
	"go-clean/model"
)

type Program string

const (
	S1D3D4 Program = "S1D3D4"
	S2     Program = "S2"
)

type Registration struct {
	ID        string  `bson:"id"`
	Name      string  `bson:"name"`
	Email     string  `bson:"email"`
	Phone     string  `bson:"phone"`
	Username  string  `bson:"username"`
	Password  string  `bson:"password"`
	Kind      Program `bson:"kind"`
	CreatedAt string  `bson:"created_at"`
}

type RegistrationRepository interface {
	Insert(register *Registration) error
	GetByEmail(email string) (*Registration, error)
	GetByPhone(phone string) (*Registration, error)
}

type RegistrationService interface {
	Create(request *model.RegistrationRequest, program Program) (*model.RegistrationResponse, error)
}
