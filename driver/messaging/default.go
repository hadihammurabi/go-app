package messaging

import (
	"errors"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver"
	"github.com/gowok/gowok/driver/messaging"
)

type def struct {
}

func ConfigureDefault(config *config.MessageBroker) (driver.Messaging, error) {
	return def{}, nil
}

func (m def) Publish(topic string, channel string, message messaging.Message) error {
	return errors.New("messaging is not available")
}

func (m def) Consume(channel string) (<-chan messaging.Message, error) {
	return nil, errors.New("messaging is not available")
}

func (m def) Ack(message messaging.Message) error {
	return errors.New("messaging is not available")
}

func (m def) IsAvailable() bool {
	return false
}
