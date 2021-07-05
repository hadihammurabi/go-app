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

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "message sent to MQ",
	})
}
