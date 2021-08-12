package main

import (
	"github.com/SemmiDev/fiber-go-clean-arch/api/infrastructure/bus"
	"github.com/SemmiDev/fiber-go-clean-arch/notifier/environments"
	"github.com/SemmiDev/fiber-go-clean-arch/notifier/mail"
	"log"
)

func main() {
	environments.New()
	dialer := mail.NewMailDialer()
	mailer := mail.NewMail(dialer)

	rabbitMQ := bus.New(environments.RabbitMQConnection)
	channel := rabbitMQ.GetChannel()
	defer channel.Close()

	channel.QueueDeclare(
		environments.RabbitMQQueue,
		true,
		false,
		false,
		false,
		nil)

	// Subscribing to QueueService1 for getting messages.
	messages, err := channel.Consume(
		environments.RabbitMQQueue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	// Build a welcome message.
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for email")
	// Make a channel to receive messages into infinite loop.
	forever := make(chan bool)

	go func() {
		for message := range messages {
			mailer.SendEmail(message.Body)
		}
	}()

	<-forever
}
