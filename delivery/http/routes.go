package http

import (
	"github.com/gofiber/fiber/v2"
)

// ConfigureRoute func
func (delivery *Delivery) ConfigureRoute() {
	delivery.HTTP.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
		})
	})

	authRouter := delivery.HTTP.Group("/auth")
	_ = NewAuth(authRouter, delivery.Service)

	usersRouter := delivery.HTTP.Group("/users")
	_ = NewUser(usersRouter, delivery.Service)
}
