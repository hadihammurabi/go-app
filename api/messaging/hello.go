package messaging

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/hadihammurabi/go-rabbitmq/exchange"
)

func Hello() error {
	conf := gowok.Get().Config.MessageBrokers["default"]
	rmq := gowok.Must(messaging.NewRabbitMQ(&conf))
	prepareHello(rmq.MQ)

	msgs := gowok.Must(rmq.Consume("hello"))
	for result := range msgs {
		fmt.Println("headers:", result.Headers)
		fmt.Println("message:", string(result.Message))
		rmq.Ack(result)
	}

	forever := make(chan bool)
	<-forever
	return nil
}

func prepareHello(mq *gorabbitmq.MQ) {
	gowok.Must(
		true,
		mq.Exchange().
			WithName("hello").
			WithType(exchange.TypeDirect).
			WithDurable(true).
			Declare(),
	)

	q := gowok.Must(mq.Queue().WithName("hello").Declare())
	gowok.Must(
		true,
		q.Binding().
			WithExchange("hello").
			Bind(),
	)
}
