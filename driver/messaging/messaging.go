package messaging

import "errors"

const (
	DriverRabbitMQ = "rabbitmq"
)

type Config struct {
	Driver string
	URL    string
}

type Messaging interface {
	Publish(topic string, channel string, message Message) error
	Consume(channel string) (<-chan Message, error)
}

func New(config Config) (Messaging, error) {
	if config.Driver == DriverRabbitMQ {
		newMq, err := ConfigureRabbitMQ(config)
		if err != nil {
			panic(err)
		}

		return newMq, nil
	}

	return nil, errors.New("messaging configuration failed")
}

type Table map[string]any

func (t Table) Validate() error {
	return nil
}

type Message struct {
	Headers     Table
	ContentType string
	Message     []byte
}
