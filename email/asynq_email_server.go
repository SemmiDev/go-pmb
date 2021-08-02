package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/tasks"
	"github.com/hibiken/asynq"
	"log"
	"os"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: os.Getenv("REDIS_DSN"), // Redis server address
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
		tasks.TypeWelcomeEmail,       // task type
		tasks.HandleEmailWelcomeTask, // handler function
	)
	// Run email server.
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
