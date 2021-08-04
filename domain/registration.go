package domain

import (
	"github.com/SemmiDev/fiber-go-clean-arch/constant"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"sync"
)

type Registration struct {
	ID             string           `bson:"id"`
	Name           string           `bson:"name"`
	Email          string           `bson:"mailer"`
	Phone          string           `bson:"phone"`
	Username       string           `bson:"username"`
	Password       string           `bson:"password"`
	Kind           constant.Program `bson:"kind"`
	Bill           constant.Bill    `bson:"bill"`
	VirtualAccount string           `bson:"virtual_account"`
	Status         bool             `bson:"status"`
	CreatedAt      string           `bson:"created_at"`
}

func NewRegistration(ID string, name string, email string, phone string, username string, password string, kind constant.Program, bill constant.Bill, virtualAccount string, status bool, createdAt string) *Registration {
	return &Registration{ID: ID, Name: name, Email: email, Phone: phone, Username: username, Password: password, Kind: kind, Bill: bill, VirtualAccount: virtualAccount, Status: status, CreatedAt: createdAt}
}

func RegisterS1D3D4PrototypePrototype() *Registration {
	return &Registration{
		Kind:   constant.S1D3D4,
		Bill:   constant.S1D3D4Bill,
		Status: false,
	}
}

func RegisterS2PrototypePrototype() *Registration {
	return &Registration{
		Kind:   constant.S2,
		Bill:   constant.S2Bill,
		Status: false,
	}
}

type RegistrationRepository interface {
	Insert(register *Registration) error
	GetByVa(va *model.UpdateStatus) (*Registration, error)
	GetByEmail(wg *sync.WaitGroup, email string)
	GetByPhone(wg *sync.WaitGroup, phone string)
	GetByUsername(username *model.LoginRequest) (*Registration, error)
	UpdateStatus(va string) error
	DeleteAll()
}

type RegistrationService interface {
	Create(request *model.RegistrationRequest, program constant.Program) (*model.RegistrationResponse, error)
	GetByUsername(req *model.LoginRequest) (*Registration, error)
	UpdateStatusBilling(va *model.UpdateStatus) error
}
