package registrant

type RegisterResponse struct {
	ID         string `json:"registrant_id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Bill       string `json:"bill"`
	PaymentURL string `json:"payment_url"`
}

type LoginResponse struct {
	AccessToken string     `json:"access_token"`
	Registrant  ReadResult `json:"registrant"`
}
