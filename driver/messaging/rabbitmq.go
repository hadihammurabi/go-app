package messaging

import (
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/streadway/amqp"
)

type rabbitmq struct {
	mq *gorabbitmq.MQ
}

func ConfigureRabbitMQ(config Config) (Messaging, error) {
	mq, err := gorabbitmq.New(config.URL)
	if err != nil {
		return nil, err
	}

	return rabbitmq{
		mq,
	}, nil
}

func (m rabbitmq) Publish(topic string, channel string, message Message) error {
	return m.mq.Publish(&gorabbitmq.MQConfigPublish{
		Exchange:   topic,
		RoutingKey: channel,
		Message: amqp.Publishing{
			Headers: amqp.Table(message.Headers),
			Body:    message.Message,
		},
	})
}

func (m rabbitmq) Consume(channel string) (<-chan Message, error) {
	q, err := m.mq.Queue().WithName(channel).Declare()
	if err != nil {
		return nil, err
	}

	result := make(chan Message)
	data, err := q.Consumer().Consume()
	if err != nil {
		return nil, err
	}

	go func() {
		for res := range data {
			result <- Message{
				Headers: Table(res.Headers),
				Message: res.Body,
			}
		}
	}()

	return result, nil
}
