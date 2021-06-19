package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-clean/internal/app/model"
	"go-clean/internal/exception"
)

func Validate(request model.CreateStudentRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.FullName, validation.Required),
		validation.Field(&request.Email, validation.Required),
		validation.Field(&request.PhoneNumber, validation.Required),
		validation.Field(&request.Path, validation.Required),
		validation.Field(&request.RegistrationNumber, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
