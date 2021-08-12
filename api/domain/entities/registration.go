package entities

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Registration struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `bson:"name"`
	Email      string             `bson:"email"`
	Phone      string             `bson:"phone"`
	Username   string             `bson:"username"`
	Password   string             `bson:"password"`
	Program    string             `bson:"kind"`
	Code       string             `bson:"code"`
	PaymentURL string             `bson:"payment_url"`
	Status     string             `bson:"status"`
	Bill       int64              `bson:"bill"`
	CreatedAt  time.Time          `bson:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at"`
}
