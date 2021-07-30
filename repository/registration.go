package repository

import (
	"go-clean/config"
	"go-clean/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const collectionName = "registrations"

type db struct {
	Collection *mongo.Collection
}

func NewRegistrationRepository(database *mongo.Database) entity.RegistrationRepository {
	return &db{
		Collection: database.Collection(collectionName),
	}
}

func (r *db) Insert(register *entity.Registration) error {
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

func (r *db) GetByEmail(email string) (register *entity.Registration, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_ = r.Collection.FindOne(ctx, bson.M{
		"email": email,
	}).Decode(&register)
	return
}

func (r *db) GetByPhone(phone string) (register *entity.Registration, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_ = r.Collection.FindOne(ctx, bson.M{
		"phone": phone,
	}).Decode(&register)
	return
}

func (r *db) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, _ = r.Collection.DeleteMany(ctx, bson.M{})
}
