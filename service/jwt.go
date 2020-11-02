package service

import (
	"belajar-go-rest-api/config"
	"belajar-go-rest-api/entities"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// JWTService struct
type JWTService struct {
	Config *config.JWTConfig
}

// NewJWTService func
func NewJWTService(jwtConfig *config.JWTConfig) entities.JWTService {
	return &JWTService{
		Config: jwtConfig,
	}
}

// Create func
func (jwtService JWTService) Create(userData *entities.User) (string, error) {
	claims := &entities.JWTClaims{}
	claims.User = userData
	claims.ExpiresAt = int64(time.Hour) * 3

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(jwtService.Config.Secret))
	if err != nil {
		return "", errors.New("Token generation fail")
	}

	return t, nil
}

// GetClaims func
func (jwtService JWTService) GetClaims(c *fiber.Ctx) (claims *entities.JWTClaims) {
	token := c.Locals("user").(*jwt.Token)
	claims = token.Claims.(*entities.JWTClaims)
	return
}
