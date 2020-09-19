package main

import (
	"belajar-go-rest-api/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=* password=* dbname=test port=5432 sslmode=disable",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

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
