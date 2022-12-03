package rest

import "github.com/gofiber/fiber/v2"

func Index(api *APIRest) {
	router := api.HTTP.Group("/auth")
	router.Get("/", api.Index)
}

func (api APIRest) Index(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Selamat datang di Belajar REST API dengan Go",
	})
}
