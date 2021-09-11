package models

import (
	"github.com/SemmiDev/go-pmb/src/registrant/entities"
	"time"
)

type QueryResult struct {
	Result interface{}
	Error  error
}

type ReadResult struct {
	ID string

	Name       string
	Email      string
	Phone      string
	Username   string
	Password   string
	Code       string
	PaymentURL string

	Program       entities.Program
	Bill          entities.Bill
	PaymentStatus entities.PaymentStatus

	CreatedDate time.Time
	LastUpdated time.Time
}
