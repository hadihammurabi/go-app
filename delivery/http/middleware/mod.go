package middleware

import "github.com/gofiber/fiber/v2"

const (
	// AUTH const
	AUTH = iota
)

// Middlewares map
var Middlewares map[int]fiber.Handler

// Use func
func Use(middlewareType int) fiber.Handler {
	return Middlewares[middlewareType]
}
