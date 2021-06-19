package config

import (
	"context"
	exception2 "go-clean/internal/exception"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

func NewMongoDatabase(configuration Config) *mongo.Database {
	ctx, cancel := NewMongoContext()
	defer cancel()

	mongoPoolMin, err := strconv.Atoi(configuration.Get("MONGO_POOL_MIN"))
	exception2.PanicIfNeeded(err)

	mongoPoolMax, err := strconv.Atoi(configuration.Get("MONGO_POOL_MAX"))
	exception2.PanicIfNeeded(err)

	mongoMaxIdleTime, err := strconv.Atoi(configuration.Get("MONGO_MAX_IDLE_TIME_SECOND"))
	exception2.PanicIfNeeded(err)

	option := options2.Client().
		SetMinPoolSize(uint64(mongoPoolMin)).
		SetMaxPoolSize(uint64(mongoPoolMax)).
		SetMaxConnIdleTime(time.Duration(mongoMaxIdleTime) * time.Second)

	client, err := mongo.NewClient(option)
	exception2.PanicIfNeeded(err)

	err = client.Connect(ctx)
	exception2.PanicIfNeeded(err)

	database := client.Database(configuration.Get("MONGO_DATABASE"))
	return database
}

func NewMongoContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10 * time.Second)
}