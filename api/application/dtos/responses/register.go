package responses

type RegisterResponse struct {
	Recipient  string `json:"recipient"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Bill       string `json:"bill"`
	PaymentURL string `bson:"payment_url"`
}
