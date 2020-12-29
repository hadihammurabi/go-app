package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// JWTClaims struct
type JWTClaims struct {
	jwt.StandardClaims
	User *User `json:"user"`
}

// JWTService interface
type JWTService interface {
	Create(userData *User) (string, error)
	GetClaims(c *fiber.Ctx) *JWTClaims
}
