package customErrors

import "errors"

var (
	InternalServerError = errors.New(InternalServerErrorMessage)
	UnprocessableEntity = errors.New(UnprocessableEntityMessage)
	BadRequest          = errors.New(BadRequestMessage)

	MissingOrMalformedJWT    = errors.New(MissingOrMalformedJWTMessage)
	InvalidTokenOrExpiredJWT = errors.New(InvalidTokenOrExpiredJWTMessage)

	UsernameIsRequired    = errors.New(UsernameIsRequiredMessage)
	PasswordIsRequired    = errors.New(PasswordIsRequiredMessage)
	NameIsRequired        = errors.New(NameIsRequiredMessage)
	EmailIsRequired       = errors.New(EmailIsRequiredMessage)
	EmailIsNotValid       = errors.New(EmailIsNotValidMessage)
	PhoneNumberIsRequired = errors.New(PhoneNumberIsRequiredMessage)
	PhoneIsNotValid       = errors.New(PhoneIsNotValidMessage)
	ProgramIsRequired     = errors.New(ProgramIsRequiredMessage)
	InvalidPassword       = errors.New(InvalidPasswordMessage)
	NotYetBIll            = errors.New(NotYetBIllMessage)

	RegisterIDIsRequired        = errors.New(RegisterIDIsRequiredMessage)
	PaymentStatusIsRequired     = errors.New(PaymentStatusIsRequiredMessage)
	FraudStatusIsRequired       = errors.New(FraudStatusIsRequiredMessage)
	PaymentTypeStatusIsRequired = errors.New(PaymentTypeStatusIsRequiredMessage)
	RegisterIDNotFound          = errors.New(RegisterIDNotFoundMessage)
	EmptyToken                  = errors.New(EmptyTokenMessage)
)
