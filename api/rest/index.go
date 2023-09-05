package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/api"
)

type index struct {
	*api.Rest
	router *fiber.App
}

func Index(r *api.Rest) *fiber.App {
	api := index{r, fiber.New()}
	api.router.Get("", api.Index)

	return api.router
}

func (api index) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Selamat datang di Belajar REST API dengan Go",
		// "token":   token,
	})
}
