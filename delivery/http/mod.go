package http

import (
	"belajar-go-rest-api/delivery/http/middleware"
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

// ConfigureRoute func
func (delivery *Delivery) ConfigureRoute() {
	delivery.HTTP.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
		})
	})

	authController := NewAuth(delivery.Service)
	auth := delivery.HTTP.Group("/auth")
	auth.Post("/login", authController.Login)
	auth.Get("/info", middleware.Auth, authController.Info)

	userController := NewUser(delivery.Service)
	users := delivery.HTTP.Group("/users")
	users.Get("/", userController.Index)
	users.Get("/:id", userController.Show)
	users.Post("/", userController.Create)
	users.Put("/:id/change-password", userController.ChangePassword)
}
