package messaging

import (
	"os"
)

type rabbitMQConfig struct {
	URL string
}

func ConfigureRabbitMQ() Messaging {
	return &rabbitMQConfig{
		URL: os.Getenv("MQ_URL"),
	}
}

func (c *rabbitMQConfig) GetURL() string {
	return c.URL
}
