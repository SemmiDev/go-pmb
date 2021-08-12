package interfaces

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IRegistrationRepository interface {
	Insert(r *entities.Registration) (register *entities.Registration, err error)
	GetByID(id primitive.ObjectID) (register *entities.Registration, err error)
	GetByUsername(username string) (register *entities.Registration, err error)
	GetByEmail(email string) (register *entities.Registration, err error)
	GetByPhone(phone string) (register *entities.Registration, err error)
	UpdateStatus(id primitive.ObjectID, status string) (err error)
	DeleteAll() (err error)
}
