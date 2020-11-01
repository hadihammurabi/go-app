package http

import (
	"belajar-go-rest-api/service"

	"github.com/gofiber/fiber/v2"
)

// Delivery struct
type Delivery struct {
	HTTP    *fiber.App
	Service *service.Service
}

// Init func
func Init(service *service.Service) *Delivery {
	app := fiber.New()

	delivery := &Delivery{
		HTTP:    app,
		Service: service,
	}
	delivery.ConfigureRoute()
	return delivery
}
