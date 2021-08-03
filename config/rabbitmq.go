package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/messaging"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
)

// ConfigureRabbitMQ func
func ConfigureRabbitMQ() (gorabbitmq.MQ, error) {
	return messaging.ConfigureRabbitMQ()
}
