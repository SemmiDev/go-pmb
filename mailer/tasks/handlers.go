package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"
)

var (
	ConfigSmtpHost     = os.Getenv("SMTPHOST")
	ConfigSmtpPort     = os.Getenv("SMTPPORT")
	ConfigSenderName   = os.Getenv("SENDERNAME")
	ConfigAuthEmail    = os.Getenv("AUTHEMAIL")
	ConfigAuthPassword = os.Getenv("AUTHPASSWORD")
)

var mailer = gomail.NewMessage()

func HandleEmailRegisterTask(ctx context.Context, t *asynq.Task) error {
	var p EmailRegisterPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Sending Email to : %s", p.Email)

	// sent email
	mailer.SetHeader("From", ConfigSenderName)
	mailer.SetHeader("To", p.Email)
	mailer.SetAddressHeader("Cc", "sammi@gmail.com", "sam")
	mailer.SetHeader("Subject", "ini credential kamu yahh")
	mailer.SetBody("text/html", "Hai kamu ❤️ "+"<br><br>"+
		"Username: "+"<b><i>"+p.Username+"</i></b>"+
		"<br>Password: "+"<b><i>"+p.Password+"</i></b>"+
		"<br>Bill: "+"<b><i>"+" Rp. "+fmt.Sprint(p.Bill)+"</i></b>"+
		"<br>VA: "+"<b><i>"+p.VirtualAccount+"</i></b>"+
		"<br>")
	//mailer.Attach("./image.jpg")

	port, err := strconv.Atoi(ConfigSmtpPort)
	dialer := gomail.NewDialer(
		ConfigSmtpHost,
		port,
		ConfigAuthEmail,
		ConfigAuthPassword,
	)

	err = dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}
	return nil
}
