package rest

import (
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/dto"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/response"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"

	"github.com/gofiber/fiber/v2"
)

type auth struct {
	*APIRest
	router *fiber.App
}

// Auth func
func Auth(r *APIRest) auth {
	api := auth{r, fiber.New()}
	api.router.Post("/login", api.Login)
	api.router.Get("/me", api.Middlewares.AuthBearer, api.Me)

	return api
}

// Login func
func (api auth) Login(c *fiber.Ctx) error {
	input := &dto.UserLoginRequest{}
	if err := c.BodyParser(input); err != nil {
		return response.Fail(c, err)
	}

	user := &entity.User{
		Email:    input.Email,
		Password: input.Password,
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
