package domain

type PaymentStatus string

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
