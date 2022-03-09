package middleware

import (
	"errors"

	jwtUtil "github.com/hadihammurabi/belajar-go-rest-api/util/jwt"

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

	// tokenData, err := m.config.Redis.Get(stringUtil.ToCacheKey("auth", "token", token))
	// if err == nil {
	// 	var user *model.User
	// 	marshalUtil.MapToStruct(tokenData.(map[string]interface{}), &user)
	// 	c.Locals("user", user)
	// 	return c.Next()
	// }

	claims, err := jwtUtil.GetJWTData(token, m.config.JWT.Secret)
	if err != nil {
		return err
	}
	user, _ := m.service.User.FindByID(c.Context(), claims.UserID)
	c.Locals("user", user)
	return c.Next()
}
