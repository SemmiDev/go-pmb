package domain

import (
	"net"
	"net/mail"
	"regexp"
	"strings"
)

type CreateRegistrantReq struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Program Program `json:"program"`
}

var phoneRegex = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

func (r *CreateRegistrantReq) Validate() error {
	if r.Name == "" {
		return RegistrantError{RegistrantErrorNameEmptyCode}
	}
	if r.Email == "" {
		return RegistrantError{RegistrantErrorEmailEmptyCode}
	}
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return RegistrantError{RegistrantErrorEmailNotValidCode}
	}
	_, err = net.LookupMX(strings.Split(r.Email, "@")[1])
	if err != nil {
		return RegistrantError{RegistrantErrorDomainNotFoundCode}
	}
	if r.Phone == "" {
		return RegistrantError{RegistrantErrorPhoneNumberEmptyCode}
	}
	if !phoneRegex.MatchString(r.Phone) {
		return RegistrantError{RegistrantErrorPhoneNumberNotValidCode}
	}
	if r.Program == "" {
		return RegistrantError{RegistrantErrorProgramEmptyCode}
	}
	if !isProgramSupported(r.Program) {
		return RegistrantError{RegistrantErrorProgramNotSupportedCode}
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
		return RegistrantError{RegistrantErrorRegistrantIdEmptyCode}
	}
	if r.PaymentStatus == "" {
		return RegistrantError{RegistrantErrorPaymentStatusEmptyCode}
	}
	if r.PaymentType == "" {
		return RegistrantError{RegistrantErrorPaymentTypeStatusEmptyCode}
	}
	if r.FraudStatus == "" {
		return RegistrantError{RegistrantErrorFraudStatusEmptyCode}
	}
	return nil
}

func isProgramSupported(program Program) bool {
	switch program {
	case ProgramS1D3D4, ProgramS2:
		return true
	default:
		return false
	}
}
