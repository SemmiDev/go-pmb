package interfaces

import (
	"github.com/streadway/amqp"
)

type IRabbitMQ interface {
	SendMessage(queue string, data interface{}) error
	GetChannel() *amqp.Channel
}
