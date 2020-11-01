package main

import (
	"belajar-go-rest-api/config/database"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_ = godotenv.Load()

	db, err := database.ConfigureDatabase()

	if err != nil {
		panic(err)
	}

	database.MigrateDatabase(db)

	app := fiber.New()

	configureRoute(app)

	app.Listen(":8080")
}
