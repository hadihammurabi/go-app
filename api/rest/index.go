package rest

import (
	"github.com/gofiber/fiber/v2"
)

type index struct {
	*APIRest
	router *fiber.App
}

func Index(r *APIRest) index {
	api := index{r, fiber.New()}
	api.router.Get("", api.Index)

	return api
}

func (api index) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Selamat datang di Belajar REST API dengan Go",
		// "token":   token,
	})
}
