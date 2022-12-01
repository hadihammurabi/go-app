package messaging

import "fmt"

func (d *APIMessaging) Hello() error {
	msgs, err := d.Config.Messaging.Consume("hello")
	if err != nil {
		return err
	}

	for result := range msgs {
		fmt.Println("headers:", result.Headers)
		fmt.Println("message:", string(result.Message))
		d.Config.Messaging.Ack(result)
	}

	forever := make(chan bool)
	<-forever

	return nil
}
