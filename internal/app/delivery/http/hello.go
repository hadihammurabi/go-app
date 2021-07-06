package http

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
	"github.com/streadway/amqp"
)

// NewHelloHandler func
func NewHelloHandler(delivery *Delivery) {
	router := delivery.HTTP.Group("/hello")
	router.Get("/", delivery.HelloIndex)
}

// HelloIndex func
func (delivery *Delivery) HelloIndex(c *fiber.Ctx) error {
	mqProcess, err := gorabbitmq.NewMQBuilder().SetConnection(&gorabbitmq.MQConfigConnection{
		URL: delivery.Config.MQ.GetURL(),
	}).SetQueue(&gorabbitmq.MQConfigQueue{
		Name: "hello:process",
	}).Build()
	if err != nil {
		return err
	}

	mqResult, err := gorabbitmq.NewMQBuilder().SetConnection(&gorabbitmq.MQConfigConnection{
		URL: delivery.Config.MQ.GetURL(),
	}).SetQueue(&gorabbitmq.MQConfigQueue{
		Name: "hello:result",
	}).Build()
	if err != nil {
		return err
	}

	defer func() {
		mqProcess.GetConnection().Close()
		mqResult.GetConnection().Close()
		mqProcess.GetChannel().Close()
		mqResult.GetChannel().Close()
	}()

	data, _ := json.Marshal(map[string]interface{}{
		"message": "hello",
	})
	err = mqProcess.Publish(&gorabbitmq.MQConfigPublish{
		RoutingKey: mqProcess.GetQueue().Name,
		Message: amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	})
	if err != nil {
		return err
	}

	mqResult.GetChannel().Qos(1, 0, false)
	msgs, err := mqResult.Consume(mqResult.GetQueue(), &gorabbitmq.MQConfigConsume{})
	if err != nil {
		return err
	}

	for msg := range msgs {
		msg.Ack(false)
		if msg.ContentType == "application/json" {
			var data map[string]interface{}
			json.Unmarshal(msg.Body, &data)
			return c.JSON(data)
		} else {
			return c.Send(msg.Body)
		}
	}

	return nil
}
