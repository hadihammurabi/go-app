package rest

import (
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/dto"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/response"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type Auth interface {
	Login(*fiber.Ctx) error
	Me(*fiber.Ctx) error
}

type auth struct {
	*APIRest
}

var _ Auth = &auth{}

// NewAuthHandler func
func NewAuthHandler(delivery *APIRest) {
	api := &auth{
		delivery,
	}
	router := api.HTTP.Group("/auth")
	router.Post("/login", api.Login)
	router.Get("/me", delivery.Middlewares.Auth, api.Me)
}

// Login func
func (api auth) Login(c *fiber.Ctx) error {
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
func (api auth) Me(c *fiber.Ctx) error {
	fromLocals := c.Locals("user").(*entity.User)
	return response.Ok(c, fromLocals)
}
