package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"go-clean/internal/app/model"
	"go-clean/internal/exception"
	"regexp"
)

func Validate(request model.CreateStudentRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.FullName, validation.Required),
		validation.Field(&request.Email, validation.Required, is.Email),
		validation.Field(&request.PhoneNumber, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{12}$"))),
		validation.Field(&request.Path, validation.Required, validation.In(uint(1), uint(2), uint(3))),
		validation.Field(&request.RegistrationNumber, validation.Required, validation.Length(10, 10)),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
