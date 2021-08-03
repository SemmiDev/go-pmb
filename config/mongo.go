package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	if err != nil {
		log.Fatalf("config.NewMongoDatabase: %v", err.Error())
	}

	mongoPoolMax, err := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	if err != nil {
		log.Fatalf("config.NewMongoDatabase: %v", err.Error())
	}

	mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	if err != nil {
		log.Fatalf("config.NewMongoDatabase: %v", err.Error())
	}

	option := options.Client().
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	if err != nil {
		log.Fatalf("config.NewMongoDatabase: %v", err.Error())
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("config.NewMongoDatabase: %v", err.Error())
	}

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
