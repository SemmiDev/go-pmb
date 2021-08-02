package config

import (
	"github.com/hibiken/asynq"
)

func NewAsynqClient(configuration Config) *asynq.Client {
	redisDSN := configuration.Get("REDIS_DSN")
	if redisDSN == "" {
		redisDSN = "127.0.0.1:6379"
	}

	asyncRedisConnection := asynq.RedisClientOpt{
		Addr: redisDSN,
	}

	return asynq.NewClient(asyncRedisConnection)
}
