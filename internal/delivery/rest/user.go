package rest

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

	uuid "github.com/satori/go.uuid"

	"github.com/gofiber/fiber/v2"
)

// NewUserHandler func
func NewUserHandler(delivery *Delivery) {
	router := delivery.HTTP.Group("/users")
	router.Get("/", delivery.UserIndex)
	router.Get("/:id", delivery.UserShow)
	router.Post("/", delivery.UserCreate)
	router.Put("/:id/change-password", delivery.UserChangePassword)
}

// UserIndex func
func (delivery *Delivery) UserIndex(c *fiber.Ctx) error {
	users, _ := delivery.Service.User.All(c.Context())
	return c.JSON(&fiber.Map{
		"data": users,
	})
}

// UserCreate func
func (delivery *Delivery) UserCreate(c *fiber.Ctx) error {
	user := &entity.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userCreated, _ := delivery.Service.User.Create(c.Context(), user)
	return c.JSON(&fiber.Map{
		"data": userCreated,
	})
}

// UserShow func
func (delivery *Delivery) UserShow(c *fiber.Ctx) error {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := delivery.Service.User.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}

// UserChangePassword func
func (delivery *Delivery) UserChangePassword(c *fiber.Ctx) error {
	id, err := uuid.FromString(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userInput := &entity.UserChangePasswordDTO{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := delivery.Service.User.ChangePassword(c.Context(), id, userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}
