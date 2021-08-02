package middleware

import (
	"errors"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
	jwtUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/jwt"
	marshalUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/marshal"
	stringUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/string"

	"github.com/gofiber/fiber/v2"
)

// Auth func
func Auth(ioc di.IOC) func(c *fiber.Ctx) error {
	config := ioc["config"].(*config.Config)
	service := ioc["service"].(*service.Service)

	return func(c *fiber.Ctx) error {
		tokenType, token, err := jwtUtil.JWTFromHeader(c.Get("Authorization"))
		if err != nil {
			return err
		}

		if tokenType != "Bearer" {
			return errors.New("invalid token")
		}

		err = config.Cache.IsAvailable()
		if config.Cache != nil && err == nil {
			tokenData, err := config.Cache.Get(stringUtil.ToCacheKey("auth", "token", token))
			if err == nil {
				var user *model.User
				marshalUtil.MapToStruct(tokenData.(map[string]interface{}), &user)
				c.Locals("user", user)
				return c.Next()
			}
		}

		claims, err := jwtUtil.GetJWTData(token, config.JWT.Secret)
		if err != nil {
			return err
		}
		user, _ := service.User.FindByID(c.Context(), claims.UserID)
		c.Locals("user", user)
		return c.Next()
	}
}
