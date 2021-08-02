package tasks

import (
	"encoding/json"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"github.com/hibiken/asynq"
)

const (
	// TypeEmailRegister is a name of the task type
	// for sending a register mailer.
	TypeEmailRegister = "mailer:register"
)

type EmailRegisterPayload struct {
	Username       string
	Password       string
	Email          string
	Bill           model.Bill
	VirtualAccount string
}

// NewRegisterEmail task payload for a new welcome mailer.
func NewRegisterEmail(username string, password string, email string, bill model.Bill, virtualAccount string) (*asynq.Task, error) {
	// Specify task payload.
	welcomePayload := EmailRegisterPayload{
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
	return asynq.NewTask(TypeEmailRegister, payload), nil
}
