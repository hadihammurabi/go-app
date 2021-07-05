package mq

import (
	"fmt"

	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/streadway/amqp"
)

func ConsumeHelloProcess(delivery *Delivery) (gorabbitmq.MQ, error) {
	mqProcess, err := gorabbitmq.NewMQBuilder().SetConnection(&gorabbitmq.MQConfigConnection{
		URL: delivery.Config.MQ.GetURL(),
	}).SetQueue(&gorabbitmq.MQConfigQueue{
		Name: "hello:process",
	}).Build()
	if err != nil {
		return nil, err
	}

	mqProcess.Consume(mqProcess.GetQueue(), &gorabbitmq.MQConfigConsume{
		OnMessage: func(msgs <-chan amqp.Delivery) {
			for msg := range msgs {
				msg.Ack(false)
				fmt.Println("GOT MESSAGE: ", string(msg.Body))
			}
		},
	})

	return mqProcess, nil
}
