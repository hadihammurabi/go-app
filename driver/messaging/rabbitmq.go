package messaging

import (
	"errors"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/messaging"
	"github.com/gowok/gowok/optional"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/streadway/amqp"
)

type rabbitmq struct {
	MQ optional.Optional[gorabbitmq.MQ]
}

func NewRabbitMQ(config *config.MessageBroker) (*rabbitmq, error) {
	mq, err := gorabbitmq.New(config.DSN)
	if err != nil {
		return nil, err
	}

	return &rabbitmq{optional.New(mq)}, nil
}

func (m rabbitmq) Publish(topic string, channel string, message messaging.Message) error {
	mq, err := m.MQ.Get()
	if err != nil {
		return errors.New("messaging is not available")
	}

	return mq.Publish(&gorabbitmq.MQConfigPublish{
		Exchange:   topic,
		RoutingKey: channel,
		Message: amqp.Publishing{
			Headers: amqp.Table(message.Headers),
			Body:    message.Message,
		},
	})
}

func (m rabbitmq) Consume(channel string) (<-chan messaging.Message, error) {
	mq, err := m.MQ.Get()
	if err != nil {
		return nil, errors.New("messaging is not available")
	}

	q, err := mq.Queue().WithName(channel).Declare()
	if err != nil {
		return nil, err
	}

	result := make(chan messaging.Message)
	data, err := q.Consumer().Consume()
	if err != nil {
		return nil, err
	}

	go func() {
		for res := range data {
			result <- messaging.Message{
				Headers: messaging.Table(res.Headers),
				Tag:     res.DeliveryTag,
				Message: res.Body,
			}
		}
	}()

	return result, nil
}

func (m rabbitmq) Ack(message messaging.Message) error {
	mq, err := m.MQ.Get()
	if err != nil {
		return errors.New("messaging is not available")
	}

	return mq.Channel().Ack(message.Tag, false)
}
