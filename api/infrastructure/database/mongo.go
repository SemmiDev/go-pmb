package database

import (
	"context"
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mongo struct {
	DB *mongo.Database
}

func NewMongoConnection() *Mongo {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin := environments.MongoPoolMin
	mongoPoolMax := environments.MongoPoolMax
	mongoMaxIdleTime := environments.MongoMaxIdleTimeSecond

	option := options.Client().
		ApplyURI(environments.MongoConnectionString).
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	database := client.Database(environments.RegistrationDatabase)
	return &Mongo{DB: database}
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
