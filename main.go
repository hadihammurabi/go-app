package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar Go REST API",
		})
	})

	app.Listen(":8080")
}
