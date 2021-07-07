package mq

import (
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/streadway/amqp"
)

func ConsumeHelloProcess(delivery *Delivery) (gorabbitmq.MQ, error) {
	mqProcess, err := gorabbitmq.NewMQBuilder().SetConnection(delivery.Config.MQ.GetURL()).
		SetQueue(&gorabbitmq.MQConfigQueue{
			Name: "hello:process",
		}).Build()
	if err != nil {
		return nil, err
	}

	mqResult, err := gorabbitmq.NewMQBuilder().SetConnection(delivery.Config.MQ.GetURL()).
		SetQueue(&gorabbitmq.MQConfigQueue{
			Name: "hello:result",
		}).Build()
	if err != nil {
		return nil, err
	}

	// defer func() {
	// 	mqProcess.Close()
	// 	mqResult.Close()
	// }()

	msgs, err := mqProcess.Consume(mqProcess.GetQueue(), &gorabbitmq.MQConfigConsume{})
	if err != nil {
		return nil, err
	}

	go func() {
		for msg := range msgs {
			msg.Ack(false)
			mqResult.Publish(&gorabbitmq.MQConfigPublish{
				RoutingKey: mqResult.GetQueue().Name,
				Message: amqp.Publishing{
					ContentType: "application/json",
					Body:        msg.Body,
				},
			})
		}
	}()

	return mqProcess, nil
}
