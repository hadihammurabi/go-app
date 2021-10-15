package rest

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/dto"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

	"github.com/gofiber/fiber/v2"
)

// NewUserHandler func
func NewUserHandler(d *APIRest) {
	router := d.HTTP.Group("/users")
	router.Get("/", d.UserIndex)
	router.Get("/:id", d.UserShow)
	router.Post("/", d.UserCreate)
	router.Put("/:id/change-password", d.UserChangePassword)
}

// UserIndex func
func (api *APIRest) UserIndex(c *fiber.Ctx) error {
	users, _ := api.Service.User.All(c.Context())
	return c.JSON(&fiber.Map{
		"data": users,
	})
}

// UserCreate func
func (api *APIRest) UserCreate(c *fiber.Ctx) error {
	user := &entity.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userCreated, _ := api.Service.User.Create(c.Context(), user)
	return c.JSON(&fiber.Map{
		"data": userCreated,
	})
}

// UserShow func
func (api *APIRest) UserShow(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := api.Service.User.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}

// UserChangePassword func
func (api *APIRest) UserChangePassword(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	userInput := &dto.UserChangePasswordRequest{}
	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	user, err := api.Service.User.ChangePassword(c.Context(), id, userInput.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.JSON(&fiber.Map{
		"data": user,
	})
}
