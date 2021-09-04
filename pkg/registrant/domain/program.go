package domain

type Program string

const (
	ProgramS1D3D4 Program = "S1D3D4"
	ProgramS2     Program = "S2"
)

func (p Program) Value() string {
	switch p {
	case ProgramS1D3D4:
		return string(ProgramS1D3D4)
	case ProgramS2:
		return string(ProgramS2)
	default:
		return ""
	}
}
