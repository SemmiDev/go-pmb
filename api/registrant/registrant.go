package registrant

import (
	"github.com/SemmiDev/go-pmb/common/helper"
	"time"
)

type Registrant struct {
	id ID

	name       string
	username   string
	password   string
	email      string
	phone      string
	code       string
	paymentURL string

	program       Program
	bill          Bill
	paymentStatus PaymentStatus

	createdDate time.Time
	lastUpdated time.Time
}

func (r Registrant) Id() ID {
	return r.id
}

func (r Registrant) Name() string {
	return r.name
}

func (r Registrant) Username() string {
	return r.username
}

func (r Registrant) Password() string {
	return r.password
}

func (r Registrant) Email() string {
	return r.email
}

func (r Registrant) Phone() string {
	return r.phone
}

func (r Registrant) Code() string {
	return r.code
}

func (r Registrant) PaymentURL() string {
	return r.paymentURL
}

func (r Registrant) Program() Program {
	return r.program
}

func (r Registrant) Bill() Bill {
	return r.bill
}

func (r Registrant) PaymentStatus() PaymentStatus {
	return r.paymentStatus
}

func (r Registrant) CreatedDate() time.Time {
	return r.createdDate
}

func (r Registrant) LastUpdated() time.Time {
	return r.lastUpdated
}

func New(name string, email string, phone string, password string, program Program) *Registrant {
	r := new(Registrant)

	now := time.Now().Local()

	registrantId := ID(helper.GenerateID())
	username := helper.GenerateUsername()
	passwordHash := helper.Hash(password)
	code := "register-registrant" + passwordHash

	r.id = registrantId
	r.name = name
	r.username = username
	r.password = passwordHash
	r.email = email
	r.phone = phone
	r.program = program
	r.code = code
	r.bill = program.Bill()
	r.paymentStatus = PaymentStatusPending
	r.createdDate = now
	r.lastUpdated = now

	return r
}
