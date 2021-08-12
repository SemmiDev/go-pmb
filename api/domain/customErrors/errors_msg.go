package customErrors

const (
	InternalServerErrorMessage = "Internal service error"
	UnprocessableEntityMessage = "Unprocessable Entity"
	BadRequestMessage          = "Bad Request"

	MissingOrMalformedJWTMessage    = "Missing or malformed JWT"
	InvalidTokenOrExpiredJWTMessage = "invalid or expired token"

	UsernameIsRequiredMessage    = "Username is required"
	PasswordIsRequiredMessage    = "Password is required"
	NameIsRequiredMessage        = "Name is required"
	EmailIsRequiredMessage       = "Email is required"
	EmailIsNotValidMessage       = "Email is not valid"
	PhoneNumberIsRequiredMessage = "Phone Number is required"
	PhoneIsNotValidMessage       = "Phone Number is not valid"
	ProgramIsRequiredMessage     = "Password is required"
	InvalidPasswordMessage       = "Password Is Invalid"
	NotYetBIllMessage            = "Please Pay The Bill First"

	RegisterIDIsRequiredMessage        = "Register ID is required"
	PaymentStatusIsRequiredMessage     = "Payment status is required"
	FraudStatusIsRequiredMessage       = "Fraud status is required"
	PaymentTypeStatusIsRequiredMessage = "Payment Type is required"

	EmailAlreadyExistsMessage = "Email already exists"
	PhoneAlreadyExistsMessage = "Phone already exists"
	RegisterIDNotFoundMessage = "Register ID not found"
	EmptyTokenMessage         = "Token is empty"
)
