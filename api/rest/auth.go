package rest

import (
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/dto"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/response"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

	"github.com/gofiber/fiber/v2"
)

// NewAuthHandler func
func NewAuthHandler(delivery *APIRest) {
	router := delivery.HTTP.Group("/auth")
	router.Post("/login", delivery.Login)
	router.Get("/me", delivery.Middlewares.Auth, delivery.Me)
}

// Login func
func (api APIRest) Login(c *fiber.Ctx) error {
	userInput := &dto.UserLoginRequest{}
	if err := c.BodyParser(userInput); err != nil {
		return response.Fail(c, err)
	}

	user := &entity.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	token, err := api.Service.Auth.Login(c.Context(), user)
	if err != nil {
		return response.Fail(c, "invalid credentials")
	}

	return response.Ok(c, &dto.UserLoginResponse{
		Token: token,
		Type:  "Bearer",
	})
}

// Me func
func (api APIRest) Me(c *fiber.Ctx) error {
	fromLocals := c.Locals("user").(*entity.User)
	return response.Ok(c, fromLocals)
}
