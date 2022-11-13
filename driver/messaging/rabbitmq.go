package messaging

import (
	"os"

	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
)

func ConfigureRabbitMQ() (*gorabbitmq.MQ, error) {
	return gorabbitmq.New(os.Getenv("MQ_URL"))
}
