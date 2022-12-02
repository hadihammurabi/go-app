package messaging

import "github.com/gowok/gowok/driver"

const (
	DriverRabbitMQ = "rabbitmq"
)

type Config struct {
	Driver string
	URL    string
}

func New(config Config) (driver.Messaging, error) {
	if config.Driver == DriverRabbitMQ {
		newMq, err := ConfigureRabbitMQ(config)
		if err != nil {
			panic(err)
		}

		return newMq, nil
	}

	return ConfigureDefault(config)
}

type Table map[string]any

func (t Table) Validate() error {
	return nil
}

type Message struct {
	Headers Table
	Tag     uint64
	Message []byte
}
