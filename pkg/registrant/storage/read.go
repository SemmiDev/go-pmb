package storage

import (
	"github.com/SemmiDev/go-pmb/pkg/registrant/entity"
	"time"
)

type ReadResult struct {
	ID string

	Name       string
	Email      string
	Phone      string
	Username   string
	Password   string
	Code       string
	PaymentURL string

	Program       entity.Program
	Bill          entity.Bill
	PaymentStatus entity.PaymentStatus

	CreatedDate time.Time
	LastUpdated time.Time
}

type QueryResult struct {
	ReadResult *ReadResult
	Error      error
}
