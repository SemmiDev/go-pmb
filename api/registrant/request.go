package registrant

import (
	"errors"
	"net"
	"net/mail"
	"regexp"
	"strings"
)

type RegisterReq struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Program Program `json:"program"`
}

var phoneRegex = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

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

type UpdatePaymentStatusReq struct {
	RegisterID    string `json:"registrant_id"`
	PaymentStatus string `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
	FraudStatus   string `json:"fraud_status"`
}

func (r *UpdatePaymentStatusReq) Validate() error {
	if r.RegisterID == "" {
		return errors.New("register id is empty")
	}
	if r.PaymentStatus == "" {
		return errors.New("payment status is empty")
	}
	if r.PaymentType == "" {
		return errors.New("payment type is empty")
	}
	if r.FraudStatus == "" {
		return errors.New("fraud status is empty")
	}

	return nil
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l *LoginReq) Validate() error {
	if l.Username == "" {
		return errors.New("username is empty")
	}
	if l.Password == "" {
		return errors.New("password is empty")
	}

	return nil
}
