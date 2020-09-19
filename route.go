package main

import (
	"github.com/gofiber/fiber/v2"
)

func configureRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
		})
	})
}
