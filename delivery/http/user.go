package http

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/service"

	uuid "github.com/satori/go.uuid"

	"github.com/gofiber/fiber/v2"
)

// UserHandler controller
type UserHandler struct {
	Service *service.Service
}

// NewUserHandler func
func NewUserHandler(router fiber.Router, service *service.Service) (userHandler *UserHandler) {
	userHandler = &UserHandler{
		Service: service,
	}

	router.Get("/", userHandler.Index)
	router.Get("/:id", userHandler.Show)
	router.Post("/", userHandler.Create)
	router.Put("/:id/change-password", userHandler.ChangePassword)

	return
}

// Index func
func (u UserHandler) Index(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"data": u.Service.User.All(),
	})
}

// Create func
func (u UserHandler) Create(c *fiber.Ctx) error {
	user := &entities.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	return c.JSON(&fiber.Map{
		"data": u.Service.User.Create(user),
	})
}

// Show func
func (u UserHandler) Show(c *fiber.Ctx) error {
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
func (u UserHandler) ChangePassword(c *fiber.Ctx) error {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userInput := &entities.UserChangePasswordDTO{}
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
