package repository

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName = "registrations"

type db struct {
	Collection *mongo.Collection
}

func NewRegistrationRepository(database *mongo.Database) model.RegistrationRepository {
	return &db{
		Collection: database.Collection(collectionName),
	}
}

func (r *db) Insert(register *model.Registration) error {
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

func (r *db) GetByVa(va *model.UpdateStatus) (account *model.Registration, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	err = r.Collection.FindOne(ctx, bson.M{"virtual_account": va.VirtualAccount}).Decode(&account)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("va not found")
		}
		return nil, err
	}
	log.Println(account)
	return account, nil
}

func (r *db) GetByEmail(email string) (account *model.Registration, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	err = r.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&account)
	if err != nil {
		return nil, err
	}
	return
}

func (r *db) GetByPhone(phone string) (account *model.Registration, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	err = r.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&account)
	if err != nil {
		return nil, err
	}
	return
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
