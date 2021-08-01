package config

import (
	"github.com/SemmiDev/fiber-go-clean-arch/auth"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	Auth   auth.AuthInterface
	Client *redis.Client
}

func NewRedisDB(host, port, password string) (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return &RedisService{
		Auth:   auth.NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
