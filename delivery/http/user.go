package http

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/service"

	uuid "github.com/satori/go.uuid"

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

// Create func
func (u User) Create(c *fiber.Ctx) error {
	user := &entities.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(&fiber.Map{
		"data": u.userService.Create(user),
	})
}

// Show func
func (u User) Show(c *fiber.Ctx) error {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := u.userService.FindByID(id)
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

	user, err := u.userService.ChangePassword(id, userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}
