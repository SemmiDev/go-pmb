package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/mailer/tasks"
	"github.com/hibiken/asynq"
	"log"
)

func main() {
	// setup configuration
	configuration := config.New()

	redisDSN := configuration.Get("REDIS_DSN")
	if redisDSN == "" {
		redisDSN = "127.0.0.1:6379"
	}

	redisConnection := asynq.RedisClientOpt{
		Addr: redisDSN, // Redis server address
	}

	worker := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: 10,
		// Specify multiple queues with different priority.
		Queues: map[string]int{
			"critical": 6, // processed 60% of the time
			"default":  3, // processed 30% of the time
			"low":      1, // processed 10% of the time
		},
	})

	// Create a new task's mux instance.
	mux := asynq.NewServeMux()
	mux.HandleFunc(
		tasks.TypeEmailRegister,       // task type
		tasks.HandleEmailRegisterTask, // handler function
	)
	// Run mailer server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
