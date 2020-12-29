package http

import (
	"belajar-go-rest-api/delivery/http/middleware"
	"belajar-go-rest-api/entity"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// NewAuthHandler func
func NewAuthHandler(delivery *Delivery) {
	router := delivery.HTTP.Group("/auth")
	router.Post("/login", delivery.Login)
	router.Get("/info", delivery.Middlewares(middleware.AUTH), delivery.Info)

	return
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

	token, err := delivery.Service.Auth.Login(user)
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
	claims := delivery.Service.JWT.GetClaims(c)
	email := claims.User.Email
	return c.SendString(fmt.Sprintf("welcome, %s", email))
}
