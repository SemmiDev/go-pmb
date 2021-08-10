package entities

import (
	"github.com/SemmiDev/fiber-go-clean-arch/internal/constant"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Registration struct {
	ID         string             `bson:"_id"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	Phone      string             `bson:"phone"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	Program    string             `bson:"kind"`
	Bill       int64              `bson:"bill"`
	Code       string             `bson:"code"`
	PaymentURL string             `bson:"payment_url"`
	Status     string             `bson:"status"`
	CreatedAt  primitive.DateTime `bson:"created_at"`
	UpdatedAt  primitive.DateTime `bson:"updated_at"`
}

var RegisterS1D3D4Prototype = &Registration{
	Program: constant.S1D3D4,
	Bill:    constant.S1D3D4Bill,
	Status:  constant.PENDING,
}

var RegisterS2Prototype = &Registration{
	Program: constant.S2,
	Bill:    constant.S2Bill,
	Status:  constant.PENDING,
}
