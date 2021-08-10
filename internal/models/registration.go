package models

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

	UpdatePaymentStatusRequest struct {
		RegisterID    string `json:"register_id"`
		PaymentStatus string `json:"payment_status"`
		PaymentType   string `json:"payment_type"`
		FraudStatus   string `json:"fraud_status"`
	}
)

func NewRegistrationResponse(recipient string, username string, password string, bill string, paymentURL string) *RegistrationResponse {
	return &RegistrationResponse{Recipient: recipient, Username: username, Password: password, Bill: bill, PaymentURL: paymentURL}
}
