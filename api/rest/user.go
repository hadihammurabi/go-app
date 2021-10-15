package rest

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest/dto"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

	"github.com/gofiber/fiber/v2"
)

type User interface {
	Index(*fiber.Ctx) error
	Store(*fiber.Ctx) error
	Show(*fiber.Ctx) error
	ChangePassword(*fiber.Ctx) error
}

type restUser struct {
	*APIRest
}

var _ User = &restUser{}

// NewUserHandler func
func NewUserHandler(d *APIRest) {
	api := &restUser{
		d,
	}
	router := d.HTTP.Group("/users")
	router.Get("/", api.Index)
	router.Get("/:id", api.Show)
	router.Post("/", api.Store)
	router.Put("/:id/change-password", api.ChangePassword)
}

// UserIndex func
func (api restUser) Index(c *fiber.Ctx) error {
	users, _ := api.Service.User.All(c.Context())
	return c.JSON(&fiber.Map{
		"data": users,
	})
}

// UserCreate func
func (api restUser) Store(c *fiber.Ctx) error {
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
func (api restUser) Show(c *fiber.Ctx) error {
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
func (api restUser) ChangePassword(c *fiber.Ctx) error {
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
