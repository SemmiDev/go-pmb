package config

import (
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Auth   auth.AuthInterface
	Client *redis.Client
}

func NewRedisDB(configuration Config) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     configuration.Get("REDIS_HOST") + ":" + configuration.Get("REDIS_PORT"),
		Password: configuration.Get("REDIS_PASSWORD"),
		DB:       0,
	})
	return &RedisService{
		Auth:   auth.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
