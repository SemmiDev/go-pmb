package repository

import (
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/SemmiDev/fiber-go-clean-arch/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RegistrationRepository interface {
	Insert(register *entity.Registration) error
	GetByID(ID string) *entity.Registration
	GetByUsername(username string) *entity.Registration
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

func (r *registrationDB) Insert(register *entity.Registration) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, register)
	if err != nil {
		return err
	}
	return nil
}

func (r *registrationDB) GetByID(ID string) *entity.Registration {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"_id": ID}).Decode(&account)
	return &account
}

func (r *registrationDB) GetByUsername(username string) *entity.Registration {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&account)

	return &account
}

func (r *registrationDB) GetByEmail(email string) bool {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&account)
	if account.ID == "" {
		return false
	}
	return true
}

func (r *registrationDB) GetByPhone(phone string) bool {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	_ = r.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&account)
	if account.ID == "" {
		return false
	}
	return true
}

func (r *registrationDB) UpdateStatus(ID string, status string) error {
	ctx, cancel := config.NewMongoContext()
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
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := r.Collection.DeleteMany(ctx, bson.M{})
	return helper.ErrorOrNil(err)
}
