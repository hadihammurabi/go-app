package http

import (
	"github.com/hadihammurabi/belajar-go-rest-api/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// ConfigureRoute func
func (delivery *Delivery) ConfigureRoute() {
	delivery.HTTP.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
		})
	})

	docs.SwaggerInfo.Host = "localhost"
	delivery.HTTP.Use("/docs", swagger.Handler)

	NewAuthHandler(delivery)
	NewUserHandler(delivery)
}
