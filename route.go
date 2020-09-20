package main

import (
	"belajar-go-rest-api/controller"

	"github.com/gofiber/fiber/v2"
)

func configureRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
		})
	})

	authController := controller.NewAuth()
	auth := app.Group("/auth")
	auth.Post("/login", authController.Login)

	userController := controller.NewUser()
	users := app.Group("/users")
	users.Get("/", userController.Index)
	users.Get("/:id", userController.Show)
	users.Post("/", userController.Create)
}
