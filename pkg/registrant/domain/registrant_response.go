package domain

type RegisterResponse struct {
	ID         ID     `json:"registrant_id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Bill       string `json:"bill"`
	PaymentURL string `json:"payment_url"`
}
