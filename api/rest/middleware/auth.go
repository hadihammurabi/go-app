package middleware

import (
	"net/http"

	jwtUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/jwt"

	"github.com/gofiber/fiber/v2"
)

// Auth func
func (m Middlewares) Auth(c *fiber.Ctx) error {
	tokenType, token, err := jwtUtil.JWTFromHeader(c.Get("Authorization"))
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("invalid token")
	}

	if tokenType != "Bearer" {
		return c.Status(http.StatusUnauthorized).SendString("invalid token")
	}

	claims, err := jwtUtil.GetJWTData(token, m.config.JWT.Secret)
	if err != nil {
		return err
	}
	user, _ := m.service.User.FindByID(c.Context(), claims.UserID)
	c.Locals("user", user)
	return c.Next()
}
