package middleware

import (
	"belajar-go-rest-api/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// NewAuthMiddleware func
func NewAuthMiddleware(jwtConfig *config.JWTConfig) func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtConfig.Secret),
	})
}
