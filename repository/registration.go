package repository

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/domain"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName = "registrations"

type db struct {
	Collection *mongo.Collection
}

func NewRegistrationRepository(database *mongo.Database) domain.RegistrationRepository {
	return &db{
		Collection: database.Collection(collectionName),
	}
}

func (r *db) Insert(register *domain.Registration) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := r.Collection.InsertOne(ctx, register)
	if err != nil {
		log.Printf("Collections: %s", collectionName)
		log.Printf("repository.Insert: %v", err.Error())
		return err
	}

	return nil
}

func (r *db) GetByVa(va *model.UpdateStatus) (*domain.Registration, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account domain.Registration
	err := r.Collection.FindOne(ctx, bson.M{"virtual_account": va.VirtualAccount}).Decode(&account)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("va not found")
		}
		return nil, err
	}
	return &account, nil
}

func (r *db) GetByUsername(req *model.LoginRequest) (*domain.Registration, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account domain.Registration
	err := r.Collection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *db) GetByEmail(email string) (*domain.Registration, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account domain.Registration
	err := r.Collection.FindOne(ctx, bson.M{"mailer": email}).Decode(&account)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &account, nil
}

func (r *db) GetByPhone(phone string) (*domain.Registration, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account domain.Registration
	err := r.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&account)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &account, nil
}

func (r *db) UpdateStatus(va string) error {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	filter := bson.M{
		"virtual_account": va,
	}

	update := bson.M{
		"$set": bson.M{
			"status": true,
		},
	}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *db) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, _ = r.Collection.DeleteMany(ctx, bson.M{})
}
