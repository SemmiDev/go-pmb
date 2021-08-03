package mailer

import (
	"fmt"
	"github.com/SemmiDev/fiber-go-clean-arch/config"
	"github.com/SemmiDev/fiber-go-clean-arch/model"
	"gopkg.in/gomail.v2"
)

type Mail struct {
	MailDialer *gomail.Dialer
	MailConfig *config.MailConfig
}

type MailService interface {
	SendEmail(template model.EmailTemplate, data interface{}) error
}

func NewMail(mailDialer *gomail.Dialer, mailConfig *config.MailConfig) MailService {
	return &Mail{MailDialer: mailDialer, MailConfig: mailConfig}
}

var mailer = gomail.NewMessage()

func (m *Mail) SendEmail(template model.EmailTemplate, data interface{}) error {
	if template == model.RegistrationTemplate {
		register := data.(*model.RegistrationResponse)

		mailer.SetHeader("From", m.MailConfig.SenderName)
		mailer.SetHeader("To", register.Recipient)
		mailer.SetAddressHeader("Cc", m.MailConfig.AddressHeaderEmail, m.MailConfig.AddressHeaderName)
		mailer.SetHeader("Subject", "ini credential kamu yahh")
		mailer.SetBody("text/html", "Hai kamu ❤️ "+"<br>"+
			"username dan password akan otomatis aktif setelah kamu membayar sesuai dengan bill ke no VA uang tertera dibawah ya <br><br>"+
			"Username: "+"<b><i>"+register.Username+"</i></b>"+
			"<br>Password: "+"<b><i>"+register.Password+"</i></b>"+
			"<br>Bill: "+"<b><i>"+" Rp. "+fmt.Sprint(register.Bill)+"</i></b>"+
			"<br>VA: "+"<b><i>"+register.VirtualAccount+"</i></b>"+
			"<br>")
		// mailer.Attach("./image.jpg")
	}

	err := m.MailDialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
