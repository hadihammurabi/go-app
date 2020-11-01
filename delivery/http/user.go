package http

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/service"

	uuid "github.com/satori/go.uuid"

	"github.com/gofiber/fiber/v2"
)

// User controller
type User struct {
	service *service.Service
}

// NewUser func
func NewUser(service *service.Service) *User {
	return &User{
		service: service,
	}
}

// Index func
func (u User) Index(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": u.service.User.All(),
	})
}

// Create func
func (u User) Create(c *fiber.Ctx) error {
	user := &entities.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(&fiber.Map{
		"data": u.service.User.Create(user),
	})
}

// Show func
func (u User) Show(c *fiber.Ctx) error {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := u.service.User.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}

// ChangePassword func
func (u User) ChangePassword(c *fiber.Ctx) error {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userInput := &struct {
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := u.service.User.ChangePassword(id, userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}
