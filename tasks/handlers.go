package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"gopkg.in/gomail.v2"
	"log"
)

const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "Universitas xxx <sammi@gmail.com>"
const CONFIG_AUTH_EMAIL = "username"
const CONFIG_AUTH_PASSWORD = "pass"

func HandleEmailWelcomeTask(ctx context.Context, t *asynq.Task) error {
	var p EmailWelcomePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		log.Println(err.Error())
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Sending Email to User to email : %s", p.Email)

	// sent
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", p.Email)
	mailer.SetAddressHeader("Cc", "sammi@gmail.com", "sam")
	mailer.SetHeader("Subject", "Ini credential kamu yah")
	mailer.SetBody("text/html", "Hai kamuu ❤❤❤️ "+"<br><br>"+
		"USERNAME: "+"<b>"+p.Username+"</b>"+
		"<br>PASSWORD: "+"<b>"+p.Password+"</b>"+
		"<br>BILL: "+"<b>"+" Rp. "+fmt.Sprint(p.Bill)+"</b>"+
		"<br>VA: "+"<b>"+p.VirtualAccount+"</b>"+
		"<br>")
	//mailer.Attach("")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)
	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
