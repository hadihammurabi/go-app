package http

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/service"

	uuid "github.com/satori/go.uuid"

	"github.com/gofiber/fiber/v2"
)

// User controller
type User struct {
	Service *service.Service
}

// NewUser func
func NewUser(router fiber.Router, service *service.Service) (user *User) {
	user = &User{
		Service: service,
	}

	router.Get("/", user.Index)
	router.Get("/:id", user.Show)
	router.Post("/", user.Create)
	router.Put("/:id/change-password", user.ChangePassword)

	return
}

// Index func
func (u User) Index(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": u.Service.User.All(),
	})
}

// Create func
func (u User) Create(c *fiber.Ctx) error {
	user := &entities.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(&fiber.Map{
		"data": u.Service.User.Create(user),
	})
}

// Show func
func (u User) Show(c *fiber.Ctx) error {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := u.Service.User.FindByID(id)
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

	user, err := u.Service.User.ChangePassword(id, userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}
