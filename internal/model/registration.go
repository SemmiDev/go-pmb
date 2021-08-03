package model

type Registration struct {
	ID             string  `bson:"id"`
	Name           string  `bson:"name"`
	Email          string  `bson:"mailer"`
	Phone          string  `bson:"phone"`
	Username       string  `bson:"username"`
	Password       string  `bson:"password"`
	Kind           Program `bson:"kind"`
	Bill           Bill    `bson:"bill"`
	VirtualAccount string  `bson:"virtual_account"`
	Status         bool    `bson:"status"`
	CreatedAt      string  `bson:"created_at"`
}

func NewRegistration(ID string, name string, email string, phone string, username string, password string, kind Program, bill Bill, virtualAccount string, status bool, createdAt string) *Registration {
	return &Registration{ID: ID, Name: name, Email: email, Phone: phone, Username: username, Password: password, Kind: kind, Bill: bill, VirtualAccount: virtualAccount, Status: status, CreatedAt: createdAt}
}

func RegisterS1D3D4PrototypePrototype() *Registration {
	return &Registration{
		Kind:   S1D3D4,
		Bill:   S1D3D4Bill,
		Status: false,
	}
}

func RegisterS2PrototypePrototype() *Registration {
	return &Registration{
		Kind:   S2,
		Bill:   S2Bill,
		Status: false,
	}
}

type RegistrationRepository interface {
	Insert(register *Registration) error
	GetByVa(va *UpdateStatus) (*Registration, error)
	GetByEmail(email string) (*Registration, error)
	GetByPhone(phone string) (*Registration, error)
	GetByUsername(username *LoginRequest) (*Registration, error)
	UpdateStatus(va string) error
	DeleteAll()
}

type RegistrationService interface {
	Create(request *RegistrationRequest, program Program) (*RegistrationResponse, error)
	GetByUsername(req *LoginRequest) (*Registration, error)
	UpdateStatusBilling(va *UpdateStatus) error
}
