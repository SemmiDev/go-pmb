package domain

import (
	"github.com/SemmiDev/go-pmb/internal/common/helper"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Registrant struct {
	RegistrantId  ID
	Name          string
	Email         string
	Phone         string
	Username      string
	Password      string
	Code          string
	PaymentURL    string
	Program       Program
	Bill          Bill
	PaymentStatus PaymentStatus
	CreatedDate   time.Time
	LastUpdated   time.Time
}

func (r *Registrant) New(name, email, phone, password string) {
	now := time.Now().Local()

	r.RegistrantId = GenerateID()
	r.Name = name
	r.Email = email
	r.Phone = phone
	r.Username = helper.GenerateUsername()
	r.Password = password
	r.HashPassword()
	r.Code = "register-registrant" + helper.Random()
	r.PaymentStatus = PaymentStatusPending
	r.CreatedDate = now
	r.LastUpdated = now
}

func (r *Registrant) ToRegistrantResponse(password string, idr string) *RegisterResponse {
	return &RegisterResponse{
		ID:         r.RegistrantId,
		Email:      r.Email,
		Username:   r.Username,
		Password:   password,
		Bill:       idr,
		PaymentURL: r.PaymentURL,
	}
}

func (r *Registrant) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	r.Password = string(bytes)
}

func (r *Registrant) IsPasswordValid(p string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(r.Password), []byte(p))
	if err != nil {
		return false, RegistrantError{RegistrantErrorWrongPasswordCode}
	}
	return true, nil
}
