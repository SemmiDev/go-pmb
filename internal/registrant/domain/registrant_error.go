package domain

// RegistrantError is a custom error from Go built-in error
type RegistrantError struct {
	Code int
}

const (
	RegistrantErrorUsernameEmptyCode = iota
	RegistrantErrorPasswordEmptyCode
	RegistrantErrorWrongPasswordCode
	RegistrantErrorNameEmptyCode
	RegistrantErrorEmailEmptyCode
	RegistrantErrorEmailNotValidCode
	RegistrantErrorPhoneNumberEmptyCode
	RegistrantErrorPhoneNumberNotValidCode
	RegistrantErrorProgramEmptyCode
	RegistrantErrorNotYetBillCode
	RegistrantErrorPaymentStatusEmptyCode
	RegistrantErrorPaymentTypeStatusEmptyCode
	RegistrantErrorFraudStatusEmptyCode
	RegistrantErrorRegisterIdNotFoundCode
	RegistrantErrorRegisterIdEmptyCode
	RegistrantErrorProgramNotSupportedCode
	RegistrantErrorDomainNotFoundCode
	RegistrantErrorRegistrantIdEmptyCode
)

func (e RegistrantError) Error() string {
	switch e.Code {
	case RegistrantErrorRegistrantIdEmptyCode:
		return "registrant id cannot be empty"
	case RegistrantErrorUsernameEmptyCode:
		return "Username cannot be empty"
	case RegistrantErrorPasswordEmptyCode:
		return "Password cannot be empty"
	case RegistrantErrorWrongPasswordCode:
		return "Wrong password"
	case RegistrantErrorNameEmptyCode:
		return "Name cannot be empty"
	case RegistrantErrorEmailEmptyCode:
		return "Email cannot be empty"
	case RegistrantErrorEmailNotValidCode:
		return "Email must be valid"
	case RegistrantErrorPhoneNumberEmptyCode:
		return "Phone number cannot be empty"
	case RegistrantErrorPhoneNumberNotValidCode:
		return "Phone number must be valid"
	case RegistrantErrorProgramEmptyCode:
		return "Program cannot be empty"
	case RegistrantErrorNotYetBillCode:
		return "Bill must be paid"
	case RegistrantErrorPaymentStatusEmptyCode:
		return "Payment status cannot be empty"
	case RegistrantErrorPaymentTypeStatusEmptyCode:
		return "Payment type status cannot be empty"
	case RegistrantErrorFraudStatusEmptyCode:
		return "Fraud status cannot be empty"
	case RegistrantErrorRegisterIdNotFoundCode:
		return "Register id not found"
	case RegistrantErrorRegisterIdEmptyCode:
		return "Register id cannot be empty"
	case RegistrantErrorProgramNotSupportedCode:
		return "Program not supported"
	case RegistrantErrorDomainNotFoundCode:
		return "Domain not found"
	default:
		return "Unrecognized user error code"
	}
}
