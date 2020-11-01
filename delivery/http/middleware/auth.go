package middleware

import "github.com/gofiber/fiber/v2"

// Auth func
func Auth(c *fiber.Ctx) error {
	return c.Next()
}
