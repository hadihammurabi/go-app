package main

import (
	"belajar-go-rest-api/model"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()


	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.User{},
	)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		users := []model.User{}
		db.Find(&users)
		return c.JSON(users)
	})

	app.Listen(":8080")
}
