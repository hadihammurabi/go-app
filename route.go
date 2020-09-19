package main

import (
	"belajar-go-rest-api/model"

	"github.com/gofiber/fiber/v2"
)

func configureRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		users := []model.User{}
		return c.JSON(users)
	})
}
