package service

import (
	"belajar-go-rest-api/config"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTService struct
type JWTService struct {
	Config *config.JWTConfig
}

// NewJWTService func
func NewJWTService() *JWTService {
	jwtConfig := config.ConfigureJWT()

	return &JWTService{
		Config: jwtConfig,
	}
}

// Create func
func (j JWTService) Create(data interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Hour * 3)
	t, err := token.SignedString([]byte(j.Config.Secret))
	if err != nil {
		return "", errors.New("Token generation fail")
	}

	return t, nil
}
