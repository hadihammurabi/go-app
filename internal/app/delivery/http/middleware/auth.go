package middleware

import (
	"errors"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/util"

	"github.com/gofiber/fiber/v2"
)

// Auth func
func Auth(config *config.Config) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenType, token, err := util.JWTFromHeader(c.Get("Authorization"))
		if err != nil {
			return err
		}

		if tokenType != "Bearer" {
			return errors.New("invalid token")
		}

		err = config.Cache.IsAvailable()
		if config.Cache != nil && err == nil {
			tokenData, err := config.Cache.Get(util.ToCacheKey("auth", "token", token))
			if err == nil {
				var User *entity.User
				util.MapToStruct(tokenData.(map[string]interface{}), &User)
				c.Locals("user", &entity.User{
					Base: entity.Base{
						ID: User.ID,
					},
				})
				return c.Next()
			}
		}

		claims, err := util.GetJWTData(token, config.JWT.Secret)
		if err != nil {
			return err
		}
		c.Locals("user", &entity.User{
			Base: entity.Base{
				ID: claims.UserID,
			},
		})
		return c.Next()
	}
}
