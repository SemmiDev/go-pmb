package config

import (
	auth2 "github.com/SemmiDev/fiber-go-clean-arch/internal/auth"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Auth   auth2.AuthInterface
	Client *redis.Client
}

func NewRedisDB(configuration Config) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     configuration.Get("REDIS_HOST") + ":" + configuration.Get("REDIS_PORT"),
		Password: configuration.Get("REDIS_PASSWORD"),
		DB:       0,
	})
	return &RedisService{
		Auth:   auth2.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
