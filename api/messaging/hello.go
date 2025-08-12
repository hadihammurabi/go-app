package messaging

import (
	"errors"
	"fmt"

	"github.com/gowok/gowok/must"
	"github.com/gowok/plugins/amqp"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/hadihammurabi/go-rabbitmq/exchange"
)

func Hello() error {
	con, ok := amqp.Connection().Get()
	if !ok {
		return errors.New("RabbitMQ not available")
	}

	mq, err := gorabbitmq.New(
		gorabbitmq.WithAMQP(con),
	)
	if err != nil {
		return err
	}

	must.Must(
		true,
		mq.Exchange().
			WithName("hello").
			WithType(exchange.TypeDirect).
			WithDurable(true).
			Declare(),
	)

	q := must.Must(mq.Queue().WithName("hello").Declare())
	must.Must(
		true,
		q.Binding().
			WithExchange("hello").
			Bind(),
	)

	msgs := must.Must(q.Consumer().Consume())
	for result := range msgs {
		fmt.Println("headers:", result.Headers)
		fmt.Println("message:", string(result.Body))
		result.Ack(false)
	}

	forever := make(chan bool)
	<-forever
	return nil
}
