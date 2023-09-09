package messaging

import (
	"errors"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver"
	"github.com/gowok/gowok/driver/messaging"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/hadihammurabi/go-rabbitmq/exchange"
	"github.com/streadway/amqp"
)

type rabbitmq struct {
	mq *gorabbitmq.MQ
}

func ConfigureRabbitMQ(config *config.MessageBroker) (driver.Messaging, error) {
	mq, err := gorabbitmq.New(config.DSN)
	if err != nil {
		return nil, err
	}

	if err = prepareHello(mq); err != nil {
		return nil, err
	}

	return rabbitmq{
		mq,
	}, nil
}

func (m rabbitmq) Publish(topic string, channel string, message messaging.Message) error {
	if !m.IsAvailable() {
		return errors.New("messaging is not available")
	}

	return m.mq.Publish(&gorabbitmq.MQConfigPublish{
		Exchange:   topic,
		RoutingKey: channel,
		Message: amqp.Publishing{
			Headers: amqp.Table(message.Headers),
			Body:    message.Message,
		},
	})
}

func (m rabbitmq) Consume(channel string) (<-chan messaging.Message, error) {
	if !m.IsAvailable() {
		return nil, errors.New("messaging is not available")
	}

	q, err := m.mq.Queue().WithName(channel).Declare()
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
	if !m.IsAvailable() {
		return errors.New("messaging is not available")
	}

	return m.mq.Channel().Ack(message.Tag, false)
}

func (m rabbitmq) IsAvailable() bool {
	return m.mq != nil
}

func prepareHello(mq *gorabbitmq.MQ) error {
	if err := mq.Exchange().WithName("hello").WithType(exchange.TypeDirect).WithDurable(true).Declare(); err != nil {
		return err
	}

	if q, err := mq.Queue().WithName("hello").Declare(); err != nil {
		return err
	} else {
		if err = q.Binding().WithExchange("hello").Bind(); err != nil {
			return err
		}
	}

	return nil
}
