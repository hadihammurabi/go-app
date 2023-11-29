package messaging

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
)

func Hello() error {
	conf := gowok.Get().Config.MessageBrokers["default"]
	rmq, err := messaging.ConfigureRabbitMQ(&conf)
	if err != nil {
		return err
	}

	msgs, err := rmq.Consume("hello")
	if err != nil {
		return err
	}

	for result := range msgs {
		fmt.Println("headers:", result.Headers)
		fmt.Println("message:", string(result.Message))
	}

	forever := make(chan bool)
	<-forever
	return nil
}
