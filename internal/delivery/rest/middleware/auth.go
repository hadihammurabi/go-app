package middleware

import (
	"errors"

	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
	jwtUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/jwt"
	marshalUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/marshal"
	stringUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/string"

	"github.com/gofiber/fiber/v2"
)

// Auth func
func (m Middlewares) Auth(c *fiber.Ctx) error {
	tokenType, token, err := jwtUtil.JWTFromHeader(c.Get("Authorization"))
	if err != nil {
		return err
	}

	if tokenType != "Bearer" {
		return errors.New("invalid token")
	}

	err = m.config.Redis.IsAvailable()
	if m.config.Redis != nil && err == nil {
		tokenData, err := m.config.Redis.Get(stringUtil.ToCacheKey("auth", "token", token))
		if err == nil {
			var user *model.User
			marshalUtil.MapToStruct(tokenData.(map[string]interface{}), &user)
			c.Locals("user", user)
			return c.Next()
		}
	}

	claims, err := jwtUtil.GetJWTData(token, m.config.JWT.Secret)
	if err != nil {
		return err
	}
	user, _ := m.service.User.FindByID(c.Context(), claims.UserID)
	c.Locals("user", user)
	return c.Next()
}
