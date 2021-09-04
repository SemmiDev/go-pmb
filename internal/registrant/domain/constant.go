package domain

type (
	Program       string
	PaymentStatus string
	Bill          int64
)

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

const (
	BillS1D3D4 Bill = 152000
	BillS2     Bill = 252000
)

func (b Bill) Value() int64 {
	switch b {
	case BillS1D3D4:
		return int64(BillS1D3D4)
	case BillS2:
		return int64(BillS2)
	default:
		return 0
	}
}

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusCancel  PaymentStatus = "cancel"
)

func (p PaymentStatus) Value() string {
	switch p {
	case PaymentStatusPending:
		return string(PaymentStatusPending)
	case PaymentStatusPaid:
		return string(PaymentStatusPaid)
	case PaymentStatusCancel:
		return string(PaymentStatusCancel)
	default:
		return ""
	}
}
