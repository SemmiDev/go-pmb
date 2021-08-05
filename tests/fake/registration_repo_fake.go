package fake

import (
	"errors"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/entity"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"sync"
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
		zap.S().Error(err.Error())
		return err
	}

	return nil
}

func (r *db) GetByVa(va *model.UpdateStatus) (*entity.Registration, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	err := r.Collection.FindOne(ctx, bson.M{"virtual_account": va.VirtualAccount}).Decode(&account)
	if err != nil {
		zap.S().Error(err.Error())
		if err.Error() == "mongo: no documents in result" {
			return nil, errors.New("va not found")
		}
		return nil, err
	}
	return &account, nil
}

func (r *db) GetByUsername(req *model.LoginRequest) (*entity.Registration, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	err := r.Collection.FindOne(ctx, bson.M{"username": req.Username}).Decode(&account)
	if err != nil {
		zap.S().Error(err.Error())
		return nil, err
	}
	return &account, nil
}

func (r *db) GetByEmail(wg *sync.WaitGroup, email string) {
	defer wg.Done()
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	// make sure not contains other error
	_ = r.Collection.FindOne(ctx, bson.M{"mailer": email}).Decode(&account)
	if account.Username != "" {
		zap.S().Error(errors.New("email was recorded"))
		Error = errors.New("email was recorded")
	}
}

func (r *db) GetByPhone(wg *sync.WaitGroup, phone string) {
	defer wg.Done()
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var account entity.Registration
	// make sure not contains other error
	_ = r.Collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&account)
	if account.Username != "" {
		zap.S().Error(errors.New("phone was recorded"))
		Error = errors.New("phone was recorded")
	}
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
		zap.S().Error(err.Error())
		return err
	}
	return nil
}

func (r *db) DeleteAll() {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	_, _ = r.Collection.DeleteMany(ctx, bson.M{})
}
