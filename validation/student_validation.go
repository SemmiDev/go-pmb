package validation

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go-clean/exception"
	"go-clean/model"
)

func Validate(request model.CreateStudentRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Identifier, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Email, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}

func ValidateProduct(request model.CreateProductRequest) {
	err := validation.ValidateStruct(&request,
		validation.Field(&request.Id, validation.Required),
		validation.Field(&request.Code, validation.Required),
		validation.Field(&request.Name, validation.Required),
		validation.Field(&request.Price, validation.Required),
		validation.Field(&request.Avaliable, validation.Required),
		validation.Field(&request.Stock, validation.Required),
	)

	if err != nil {
		panic(exception.ValidationError{
			Message: err.Error(),
		})
	}
}
