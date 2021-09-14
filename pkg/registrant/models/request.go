package models

import (
	"errors"
	"github.com/SemmiDev/go-pmb/pkg/registrant/entity"
	"net"
	"net/mail"
	"regexp"
	"strings"
)

// RegisterReq for Register request payload.
type RegisterReq struct {
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Phone   string         `json:"phone"`
	Program entity.Program `json:"program"`
}

// phoneRegex for rules of phone number.
var phoneRegex = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

// Validate for validate RegisterReq.
func (r *RegisterReq) Validate() error {
	if r.Name == "" {
		return errors.New("name is empty")
	}
	if r.Email == "" {
		return errors.New("email is empty")
	}
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	_, err = net.LookupMX(strings.Split(r.Email, "@")[1])
	if err != nil {
		return errors.New("invalid email")
	}
	if r.Phone == "" {
		return errors.New("phone is empty")
	}
	if !phoneRegex.MatchString(r.Phone) {
		return errors.New("invalid phone number")
	}
	if r.Program.Empty() {
		return errors.New("program is empty")
	}
	if !r.Program.IsSupported() {
		return errors.New("program is not supported")
	}

	return nil
}

// UpdatePaymentStatusReq for update payments status request payload.
type UpdatePaymentStatusReq struct {
	RegisterID    string `json:"registrant_id"`
	PaymentStatus string `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
	FraudStatus   string `json:"fraud_status"`
}

// Validate for validate UpdatePaymentStatusReq.
func (r *UpdatePaymentStatusReq) Validate() error {
	if r.RegisterID == "" {
		return errors.New("register id is empty")
	}
	if r.PaymentStatus == "" {
		return errors.New("payments status is empty")
	}
	if r.PaymentType == "" {
		return errors.New("payments type is empty")
	}
	if r.FraudStatus == "" {
		return errors.New("fraud status is empty")
	}

	return nil
}

// LoginReq for login request payload.
type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate for validate LoginReq.
func (l *LoginReq) Validate() error {
	if l.Username == "" {
		return errors.New("username is empty")
	}
	if l.Password == "" {
		return errors.New("password is empty")
	}

	return nil
}
