package models

// RegisterResponse for register response payload.
type RegisterResponse struct {
	ID         string `json:"registrant_id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Bill       string `json:"bill"`
	PaymentURL string `json:"payment_url"`
}

// LoginResponse for login response payload.
type LoginResponse struct {
	AccessToken string     `json:"access_token"`
	Registrant  ReadResult `json:"registrant"`
}
