package mail

import (
	"encoding/json"
	"github.com/SemmiDev/fiber-go-clean-arch/notifier/environments"
	"gopkg.in/gomail.v2"
	"log"
)

type IEmail interface {
	SendEmail(data []byte) error
}

type Email struct {
	MailDialer *gomail.Dialer
}

func NewMail(mailDialer *gomail.Dialer) IEmail {
	return &Email{MailDialer: mailDialer}
}

var mailer = gomail.NewMessage()

func (e *Email) SendEmail(data []byte) error {
	var response RegisterResponse
	err := json.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	mailer.SetHeader("From", environments.EmailSenderName)
	mailer.SetHeader("To", response.Recipient)
	mailer.SetAddressHeader("Cc", environments.EmailAddressHeaderEmail, environments.EmailAddressHeaderName)
	mailer.SetHeader("Subject", "ini credential kamu yahh")
	mailer.SetBody("text/html", "Hai kamu ❤️ "+"<br>"+
		"bayar dlu pake payments url dibawah ini agar username dan password kamu aktif <br><br>"+
		"Username: "+"<b><i>"+response.Username+"</i></b>"+
		"<br>Password: "+"<b><i>"+response.Password+"</i></b>"+
		"<br>Bill: "+"<b><i>"+""+response.Bill+"</i></b>"+
		"<br>Payment URL: "+"<b><i>"+response.PaymentURL+"</i></b>"+
		"<br>")
	// mailer.Attach("./image.jpg")

	err = e.MailDialer.DialAndSend(mailer)
	if err != nil {
		panic(err)
	}

	log.Printf("-> email has been processed: %s\n", response.Recipient)
	return nil
}
