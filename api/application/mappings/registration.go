package mappings

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/constants"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func NewRegistrationTemplate(name string, email string, username string, password string, phone string, program string) *entities.Registration {
	if program == "S1D3D4" {
		return &entities.Registration{
			ID:        primitive.NewObjectID(),
			Program:   constants.S1D3D4,
			Bill:      constants.S1D3D4Bill,
			Status:    constants.PaymentStatusPending,
			Name:      name,
			Username:  username,
			Password:  password,
			Email:     email,
			Phone:     phone,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	}
	return &entities.Registration{
		ID:        primitive.NewObjectID(),
		Program:   constants.S2,
		Bill:      constants.S2Bill,
		Status:    constants.PaymentStatusPending,
		Name:      name,
		Username:  username,
		Password:  password,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
