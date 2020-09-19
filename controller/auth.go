package controller

import (
	"belajar-go-rest-api/model"
	"belajar-go-rest-api/service"

	"github.com/gofiber/fiber/v2"
)

// Auth controller
type Auth struct {
	authService *service.Auth
}

type user struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// NewAuth func
func NewAuth() *Auth {
	return &Auth{
		authService: service.NewAuth(),
	}
}

// Login func
func (a Auth) Login(c *fiber.Ctx) error {
	userInput := &user{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := &model.User{
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	token, err := a.authService.Login(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(&fiber.Map{
		"token": token,
	})
}
