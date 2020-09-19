package main

import (
	"belajar-go-rest-api/config"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_ = godotenv.Load()

	db, err := config.ConfigureDatabase()

	if err != nil {
		panic(err)
	}

	config.MigrateDatabase(db)

	app := fiber.New()

	configureRoute(app)

	app.Listen(":8080")
}
