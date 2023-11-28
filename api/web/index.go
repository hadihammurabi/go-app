package web

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/exception"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web/dto"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/api"
)

type index struct {
	*api.Rest
	router *fiber.App
	redis  *redis.Client
}

func Index(r *api.Rest) *fiber.App {
	rdb := gowok.Get().Redis("cache").OrPanic(exception.ErrNoDatabaseFound)
	api := index{r, fiber.New(), rdb}
	api.router.Get("", api.Index)

	return api.router
}

func (api index) Index(c *fiber.Ctx) error {
	token, err := api.redis.Get(c.Context(), "token").Result()
	if err != nil && err != redis.Nil {
		return dto.Fail(c, err)
	}

	if token == "" {
		err := api.redis.Set(
			c.Context(),
			"token",
			"adacd8a852a0813c9bf8e7690f4461d56930c867e241c55eac0afa5d7dd9ac87",
			time.Hour,
		).Err()
		if err != nil {
			return dto.Fail(c, err)
		}
	}
	return c.JSON(fiber.Map{
		"message": "Selamat datang di Belajar REST API dengan Go",
		"token":   token,
	})
}
