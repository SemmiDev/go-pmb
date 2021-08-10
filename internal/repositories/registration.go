package repositories

import (
	"github.com/SemmiDev/fiber-go-clean-arch/internal/entities"
	"github.com/SemmiDev/fiber-go-clean-arch/internal/helper"
	config2 "github.com/SemmiDev/fiber-go-clean-arch/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RegistrationRepository interface {
	Insert(register *entities.Registration) error
	GetByID(ID string) *entities.Registration
	GetByUsername(username string) *entities.Registration
	GetByEmail(email string) bool
	GetByPhone(phone string) bool
	UpdateStatus(ID string, status string) error
	DeleteAll() error
}

type registrationDB struct {
	Collection *mongo.Collection
}

func NewRegistrationRepository(database *mongo.Database) RegistrationRepository {
	return &registrationDB{
		Collection: database.Collection("registrations"),
	}
}

func (r *registrationDB) Insert(register *entities.Registration) error {
	ctx, cancel := config2.NewMongoContext()
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, register)
	if err != nil {
		return err
	}
	return nil
}

func (r *registrationDB) GetByID(ID string) *entities.Registration {
	ctx, cancel := config2.NewMongoContext()
	defer cancel()

	var account entities.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&account)
	return &account
}

func (r *registrationDB) GetByUsername(username string) *entities.Registration {
	ctx, cancel := config2.NewMongoContext()
	defer cancel()

	var account entities.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&account)

	return &account
}

func (r *registrationDB) GetByEmail(email string) bool {
	ctx, cancel := config2.NewMongoContext()
	defer cancel()

	var account entities.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&account)
	if account.ID == "" {
		return false
	}
	return true
}

func (r *registrationDB) GetByPhone(phone string) bool {
	ctx, cancel := config2.NewMongoContext()
	defer cancel()

	var account entities.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&account)
	if account.ID == "" {
		return false
	}
	return true
}

func (r *registrationDB) UpdateStatus(ID string, status string) error {
	ctx, cancel := config2.NewMongoContext()
	defer cancel()

	filter := bson.M{
		"_id": ID,
	}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": primitive.NewDateTimeFromTime(time.Now()),
		},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return helper.ErrorOrNil(err)
}

func (r *registrationDB) DeleteAll() error {
	ctx, cancel := config2.NewMongoContext()
	defer cancel()

	_, err := r.Collection.DeleteMany(ctx, bson.M{})
	return helper.ErrorOrNil(err)
}
