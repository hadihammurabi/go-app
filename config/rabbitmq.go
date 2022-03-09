package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
)

// ConfigureRabbitMQ func
func ConfigureRabbitMQ() (gorabbitmq.MQ, error) {
	return messaging.ConfigureRabbitMQ()
}
