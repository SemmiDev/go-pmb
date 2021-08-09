package model

type (
	RegistrationRequest struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Program string `json:"program"`
	}

	RegistrationResponse struct {
		Recipient  string `json:"recipient"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Bill       string `json:"bill"`
		PaymentURL string `bson:"payment_url"`
	}

	UpdatePaymentStatus struct {
		RegisterID    string `json:"register_id"`
		PaymentStatus string `json:"payment_status"`
		PaymentType   string `json:"payment_type"`
		FraudStatus   string `json:"fraud_status"`
	}
)
