package constant

type (
	Program       string
	Bill          int64
	EmailTemplate int
)

const (
	S1D3D4 Program = "S1D3D4"
	S2     Program = "S2"

	S1D3D4Bill Bill = 152000
	S2Bill     Bill = 252000

	RegistrationTemplate EmailTemplate = iota
	ForgotPasswordTemplate
)
