package bus

import (
	"encoding/json"
	"github.com/SemmiDev/fiber-go-clean-arch/api/application/dtos/responses"
	"github.com/SemmiDev/fiber-go-clean-arch/api/domain/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func New(connectionString string) interfaces.IRabbitMQ {
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		log.Fatal(err)
	}

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err)
	}

	return &RabbitMQ{Connection: connection, Channel: channel}
}

func (r *RabbitMQ) SendMessage(queue string, data interface{}) error {
	_, err := r.Channel.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		return err
	}

	switch queue {
	case "registration-notifier":
		payload, err := json.Marshal(data.(*responses.RegisterResponse))
		if err != nil {
			return err
		}

		if err := r.Channel.Publish("", queue, false, false, amqp.Publishing{
			ContentType: fiber.MIMEApplicationJSON,
			Body:        payload,
		}); err != nil {
			return err
		}
	case "other":
		payload := data.(string)
		if err := r.Channel.Publish("", queue, false, false, amqp.Publishing{
			ContentType: fiber.MIMETextPlain,
			Body:        []byte(payload),
		}); err != nil {
			return err
		}
	}

	return nil
}

func (r *RabbitMQ) GetChannel() *amqp.Channel {
	return r.Channel
}
