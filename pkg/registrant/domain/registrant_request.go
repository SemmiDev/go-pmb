package domain

import (
	"github.com/SemmiDev/go-pmb/pkg/registrant/errors"
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
		return errors.RegistrantError{errors.RegistrantErrorNameEmptyCode}
	}
	if r.Email == "" {
		return errors.RegistrantError{errors.RegistrantErrorEmailEmptyCode}
	}
	_, err := mail.ParseAddress(r.Email)
	if err != nil {
		return errors.RegistrantError{errors.RegistrantErrorEmailNotValidCode}
	}
	_, err = net.LookupMX(strings.Split(r.Email, "@")[1])
	if err != nil {
		return errors.RegistrantError{errors.RegistrantErrorDomainNotFoundCode}
	}
	if r.Phone == "" {
		return errors.RegistrantError{errors.RegistrantErrorPhoneNumberEmptyCode}
	}
	if !phoneRegex.MatchString(r.Phone) {
		return errors.RegistrantError{errors.RegistrantErrorPhoneNumberNotValidCode}
	}
	if r.Program == "" {
		return errors.RegistrantError{errors.RegistrantErrorProgramEmptyCode}
	}
	if !isProgramSupported(r.Program) {
		return errors.RegistrantError{errors.RegistrantErrorProgramNotSupportedCode}
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
		return errors.RegistrantError{errors.RegistrantErrorRegistrantIdEmptyCode}
	}
	if r.PaymentStatus == "" {
		return errors.RegistrantError{errors.RegistrantErrorPaymentStatusEmptyCode}
	}
	if r.PaymentType == "" {
		return errors.RegistrantError{errors.RegistrantErrorPaymentTypeStatusEmptyCode}
	}
	if r.FraudStatus == "" {
		return errors.RegistrantError{errors.RegistrantErrorFraudStatusEmptyCode}
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
