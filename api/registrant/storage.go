package registrant

import (
	"time"
)

type ReadResult struct {
	RegistrantId ID

	Name       string
	Email      string
	Phone      string
	Username   string
	Password   string
	Code       string
	PaymentURL string

	Program       Program
	Bill          Bill
	PaymentStatus PaymentStatus

	CreatedDate time.Time
	LastUpdated time.Time
}
