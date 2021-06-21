package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"os"
)

func NewRedisClient() *redis.Client {
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return client
}
