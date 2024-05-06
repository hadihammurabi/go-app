package web

import (
	"github.com/eko/gocache/lib/v4/cache"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web/dto"
)

func Index(rdb *cache.Cache[any]) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token, err := rdb.Get(c.Context(), "token")
		if err != nil && err != redis.Nil {
			return dto.Fail(c, err)
		}

		if token == "" {
			err := rdb.Set(
				c.Context(),
				"token",
				"adacd8a852a0813c9bf8e7690f4461d56930c867e241c55eac0afa5d7dd9ac87",
			)
			if err != nil {
				return dto.Fail(c, err)
			}
		}

		return c.JSON(fiber.Map{
			"message": "Selamat datang di Belajar REST API dengan Go",
			"token":   token,
		})
	}
}
