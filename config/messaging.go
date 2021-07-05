package config

import (
	"errors"
	"os"

	"github.com/hadihammurabi/belajar-go-rest-api/platform/messaging"
)

const (
	driverRabbitMQ = "rabbitmq"
)

// ConfigureMQ func
func ConfigureMQ() (messaging.Messaging, error) {
	driver := os.Getenv("MQ_DRIVER")
	if driver == "" {
		driver = driverRabbitMQ
	}

	if driver == driverRabbitMQ {
		mqConfig := messaging.ConfigureRabbitMQ()
		return mqConfig, nil
	}

	return nil, errors.New("unknown cache driver")
}
