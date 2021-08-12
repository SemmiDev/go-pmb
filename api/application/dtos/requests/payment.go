package requests

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/customErrors"
)

type UpdatePaymentStatus struct {
	RegisterID    string `json:"register_id"`
	PaymentStatus string `json:"payment_status"`
	PaymentType   string `json:"payment_type"`
	FraudStatus   string `json:"fraud_status"`
}

func (r *UpdatePaymentStatus) IsValid() error {
	if r.RegisterID == "" {
		return customErrors.RegisterIDIsRequired
	}
	if r.PaymentStatus == "" {
		return customErrors.PaymentStatusIsRequired
	}
	if r.FraudStatus == "" {
		return customErrors.FraudStatusIsRequired
	}
	if r.PaymentType == "" {
		return customErrors.PaymentTypeStatusIsRequired
	}

	return nil
}
