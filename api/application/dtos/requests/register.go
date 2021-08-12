package requests

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/customErrors"
	"regexp"
)

type Register struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Program string `json:"program"`
}

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	phoneRegexp = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
)

func (r *Register) IsValid() error {
	if r.Name == "" {
		return customErrors.NameIsRequired
	}

	if r.Email == "" {
		return customErrors.EmailIsRequired
	}

	if r.Phone == "" {
		return customErrors.PhoneNumberIsRequired
	}

	if r.Program == "" {
		return customErrors.ProgramIsRequired
	}

	if emailRegexp.MatchString(r.Email) == false {
		return customErrors.EmailIsNotValid
	}

	if phoneRegexp.MatchString(r.Phone) == false {
		return customErrors.PhoneIsNotValid
	}

	return nil
}
