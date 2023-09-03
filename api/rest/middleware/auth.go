package middleware

import (
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
)

// AuthBearer func
func (m Middlewares) AuthBearer(c *fiber.Ctx) error {
	bearerHeader := c.Get("Authorization")
	bearerSplit := strings.Split(bearerHeader, " ")
	if len(bearerSplit) != 2 {
		return c.Status(http.StatusUnauthorized).SendString("invalid token")
	}

	tokenType, token := bearerSplit[0], bearerSplit[1]
	if tokenType != "Bearer" {
		return c.Status(http.StatusUnauthorized).SendString("invalid token")
	}

	conf := driver.Get().Config
	parsed, err := jwt.ParseWithClaims(token, &entity.JWTClaims{}, func(t *jwt.Token) (any, error) {
		return []byte(conf.Security.Secret), nil
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).SendString("invalid token")
	}

	claims, ok := parsed.Claims.(*entity.JWTClaims)
	if !ok || !parsed.Valid {
		return c.Status(http.StatusUnauthorized).SendString("invalid token")
	}

	user, _ := m.service.User.FindByID(c.Context(), claims.UserID)
	c.Locals("user", user)
	return c.Next()
}
