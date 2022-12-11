package rest

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
)

func Index(api *APIRest) {
	router := api.HTTP.Group("/")
	router.Get("", api.Index)
}

func (api APIRest) Index(c *fiber.Ctx) error {
	redis := ioc.Get(database.Redis{})
	token, err := redis.Get(c.Context(), "token").Result()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(fiber.Map{
		"message": "Selamat datang di Belajar REST API dengan Go",
		"token":   token,
	})
}
