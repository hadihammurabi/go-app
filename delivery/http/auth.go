package http

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
func NewAuth(database *gorm.DB) *Auth {
	return &Auth{
		authService: service.NewAuth(database),
	}
}

// Login func
func (a Auth) Login(c *fiber.Ctx) error {
	userInput := &user{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user := &entities.User{
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

// Info func
func (a Auth) Info(c *fiber.Ctx) error {
	return c.SendString("not implemented yet")
}
