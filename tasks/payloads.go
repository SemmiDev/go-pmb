package tasks

import (
	"encoding/json"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/hibiken/asynq"
)

const (
	// TypeWelcomeEmail is a name of the task type
	// for sending a welcome email.
	TypeWelcomeEmail = "email:welcome"
)

type EmailWelcomePayload struct {
	Username       string
	Password       string
	Email          string
	Bill           model.Bill
	VirtualAccount string
}

// NewWelcomeEmailTask task payload for a new welcome email.
func NewWelcomeEmailTask(username string, password string, email string, bill model.Bill, virtualAccount string) (*asynq.Task, error) {
	// Specify task payload.
	welcomePayload := EmailWelcomePayload{
		Username:       username,
		Password:       password,
		Email:          email,
		Bill:           bill,
		VirtualAccount: virtualAccount,
	}

	payload, err := json.Marshal(welcomePayload)
	if err != nil {
		return nil, err
	}

	// Return a new task with given type and payload.
	return asynq.NewTask(TypeWelcomeEmail, payload), nil
}
