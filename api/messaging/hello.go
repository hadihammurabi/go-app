package messaging

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/messaging"
)

type hello driver.Driver

func Hello() {
	m := hello(*driver.Get())
	gowok.Must(0, m.Hello())
}

func (d *hello) Hello() error {
	conf := d.Config.MessageBrokers[""]
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
