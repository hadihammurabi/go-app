package controller

import (
	"belajar-go-rest-api/service"

	"github.com/gofiber/fiber/v2"
)

// User controller
type User struct {
	userService *service.User
}

// NewUser func
func NewUser() *User {
	return &User{
		userService: service.NewUser(),
	}
}

// Index func
func (u User) Index(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": u.userService.All(),
	})
}
