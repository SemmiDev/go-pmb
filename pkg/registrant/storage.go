package registrant

import "time"

type ReadResult struct {
	ID string

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

type QueryResult struct {
	ReadResult *ReadResult
	Error      error
}
