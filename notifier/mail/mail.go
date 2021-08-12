package mail

import (
	"github.com/SemmiDev/fiber-go-clean-arch/notifier/environments"
	"gopkg.in/gomail.v2"
)

type RegisterResponse struct {
	Recipient  string `json:"recipient"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Bill       string `json:"bill"`
	PaymentURL string `bson:"payment_url"`
}

func NewMailDialer() *gomail.Dialer {
	return gomail.NewDialer(
		environments.EmailSmtpHost,
		environments.EmailSmtpPort,
		environments.EmailAuth,
		environments.EmailPassword,
	)
}
