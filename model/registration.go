package model

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

var RegisterS1D3D4Prototype = &Registration{
	Kind:          S1D3D4,
	Bill:          S1D3D4Bill,
	AccountNumber: S1D3D4AccountNumber,
	Status:        false,
}

var RegisterS2Prototype = &Registration{
	Kind:          S2,
	Bill:          S2Bill,
	AccountNumber: S2AccountNumber,
	Status:        false,
}

type RegistrationRepository interface {
	Insert(register *Registration) error
	GetByEmail(email string) (*Registration, error)
	GetByPhone(phone string) (*Registration, error)
	DeleteAll()
}

type RegistrationService interface {
	Create(request *RegistrationRequest, program Program) (*RegistrationResponse, error)
}
