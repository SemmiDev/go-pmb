package entities

import (
	"github.com/SemmiDev/go-pmb/src/common/helper"
	"github.com/google/uuid"
	"time"
)

type Registrant struct {
	ID string

	Name       string
	Username   string
	Password   string
	Email      string
	Phone      string
	Code       string
	PaymentURL string

	Program       Program
	Bill          Bill
	PaymentStatus PaymentStatus

	CreatedDate time.Time
	LastUpdated time.Time
}

func NewRegistrant(name, email, phone, password string, program Program) *Registrant {
	r := new(Registrant)
	now := time.Now().Local()

	r.ID = uuid.NewString()
	r.Username = helper.GenerateUsername()
	r.Password = helper.Hash(password)
	r.Email = email
	r.Phone = phone
	r.Program = program
	r.Code = "register-registrant" + r.Username + r.Password
	r.Bill = program.Bill()
	r.PaymentStatus = PaymentStatusPending
	r.CreatedDate = now
	r.LastUpdated = now

	return r
}
