package http

import (
	"belajar-go-rest-api/delivery/http/middleware"
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/service"

	"github.com/gofiber/fiber/v2"
)

// AuthHandler controller
type AuthHandler struct {
	Service *service.Service
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewAuthHandler func
func NewAuthHandler(router fiber.Router, service *service.Service) (authHandler *AuthHandler) {
	authHandler = &AuthHandler{
		Service: service,
	}

	router.Post("/login", authHandler.Login)
	router.Get("/info", middleware.Auth, authHandler.Info)

	return
}

// Login func
func (a AuthHandler) Login(c *fiber.Ctx) error {
	userInput := &user{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := &entities.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	token, err := a.Service.Auth.Login(user)
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
func (a AuthHandler) Info(c *fiber.Ctx) error {
	return c.SendString("not implemented yet")
}
