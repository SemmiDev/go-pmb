package database

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/environments"
	"github.com/go-redis/redis/v8"
)

type IRedisConnection interface {
	Disconnect()
}

type RedisConnection struct {
	Client *redis.Client
}

func (r *RedisConnection) Disconnect() {
	r.Client.Close()
}

func NewRedisConnection() *RedisConnection {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     environments.RedisConnectionString,
		Password: environments.RedisPassword,
		DB:       0,
	})

	return &RedisConnection{
		Client: redisClient,
	}
}
