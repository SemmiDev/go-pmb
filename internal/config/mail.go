package config

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

type MailConfig struct {
	SmtpHost           string
	SmtpPort           int
	SenderName         string
	AuthEmail          string
	AuthPassword       string
	AddressHeaderEmail string
	AddressHeaderName  string
}

func NewMailDialer(configuration Config) (*gomail.Dialer, *MailConfig) {
	smtpPort := configuration.Get("SMTP_PORT")
	port, _ := strconv.Atoi(smtpPort)
	mailConfig := MailConfig{
		SmtpHost:           configuration.Get("SMTP_HOST"),
		SmtpPort:           port,
		SenderName:         configuration.Get("SENDER_NAME"),
		AuthEmail:          configuration.Get("AUTH_EMAIL"),
		AuthPassword:       configuration.Get("AUTH_PASSWORD"),
		AddressHeaderEmail: configuration.Get("ADDRESS_HEADER_EMAIL"),
		AddressHeaderName:  configuration.Get("ADDRESS_HEADER_NAME"),
	}

	dialer := gomail.NewDialer(
		mailConfig.SmtpHost,
		mailConfig.SmtpPort,
		mailConfig.AuthEmail,
		mailConfig.AuthPassword,
	)
	return dialer, &mailConfig
}
