package repositories

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/entities"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/database"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RegistrationRepository struct {
	Collection *mongo.Collection
}

func NewRegistrationRepository(database *database.Mongo) interfaces.IRegistrationRepository {
	return &RegistrationRepository{
		Collection: database.DB.Collection(environments.RegistrationCollection),
	}
}

func (rr *RegistrationRepository) Insert(r *entities.Registration) (register *entities.Registration, err error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	_, err = rr.Collection.InsertOne(ctx, &r)
	if err != nil {
		return nil, errors.New("failed to register!")
	}

	return r, nil
}

func (rr *RegistrationRepository) GetByID(id primitive.ObjectID) (register *entities.Registration, err error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	if err := rr.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&register); err != nil {
		return nil, nil
	}

	return register, nil
}

func (rr *RegistrationRepository) GetByUsername(username string) (register *entities.Registration, err error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	if err := rr.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&register); err != nil {
		return nil, nil
	}

	return register, nil
}

func (rr *RegistrationRepository) GetByEmail(email string) (register *entities.Registration, err error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	if err := rr.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&register); err != nil {
		return nil, nil
	}

	return register, nil
}

func (rr RegistrationRepository) GetByPhone(phone string) (register *entities.Registration, err error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	if err := rr.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&register); err != nil {
		return nil, nil
	}

	return register, nil
}

func (rr RegistrationRepository) UpdateStatus(id primitive.ObjectID, status string) (err error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	filter := bson.M{
		"_id": id,
	}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}

	_, err = rr.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (rr RegistrationRepository) DeleteAll() (err error) {
	ctx, cancel := database.NewMongoContext()
	defer cancel()

	_, err = rr.Collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return err
	}

	return nil
}
