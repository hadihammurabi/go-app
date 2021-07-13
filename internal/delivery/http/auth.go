package http

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/http/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

	"github.com/gofiber/fiber/v2"
)

// NewAuthHandler func
func NewAuthHandler(delivery *Delivery) {
	router := delivery.HTTP.Group("/auth")
	router.Post("/login", delivery.Login)
	router.Get("/info", delivery.Middlewares(middleware.AUTH), delivery.Info)
}

// Login func
func (delivery Delivery) Login(c *fiber.Ctx) error {
	userInput := &entity.UserLoginDTO{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := &entity.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	token, err := delivery.Service.Auth.Login(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"token": token,
	})
}

// Info func
func (delivery Delivery) Info(c *fiber.Ctx) error {
	fromLocals := c.Locals("user").(*entity.User)
	return c.JSON(fromLocals)
}
