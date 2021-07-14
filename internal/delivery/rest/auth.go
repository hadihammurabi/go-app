package rest

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest/middleware"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest/response"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/dto"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"

	"github.com/gofiber/fiber/v2"
)

// NewAuthHandler func
func NewAuthHandler(delivery *Delivery) {
	router := delivery.HTTP.Group("/auth")
	router.Post("/login", delivery.Login)
	router.Get("/me", delivery.Middlewares(middleware.AUTH), delivery.Me)
}

// Login func
func (delivery Delivery) Login(c *fiber.Ctx) error {
	userInput := &dto.UserLoginRequest{}
	if err := c.BodyParser(userInput); err != nil {
		return response.Fail(c, err)
	}

	user := &model.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	token, err := delivery.Service.Auth.Login(c.Context(), user)
	if err != nil {
		return response.Fail(c, "invalid credentials")
	}

	return response.Ok(c, &fiber.Map{
		"token": token,
		"type":  "Bearer",
	})
}

// Me func
func (delivery Delivery) Me(c *fiber.Ctx) error {
	fromLocals := c.Locals("user").(*model.User)
	return response.Ok(c, fromLocals)
}
