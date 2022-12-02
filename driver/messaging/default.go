package messaging

import (
	"errors"
)

type def struct {
}

func ConfigureDefault(config Config) (Messaging, error) {
	return def{}, nil
}

func (m def) Publish(topic string, channel string, message Message) error {
	return errors.New("messaging is not available")
}

func (m def) Consume(channel string) (<-chan Message, error) {
	return nil, errors.New("messaging is not available")
}

func (m def) Ack(message Message) error {
	return errors.New("messaging is not available")
}

func (m def) IsAvailable() bool {
	return false
}
