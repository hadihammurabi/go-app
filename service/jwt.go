package service

import (
	"errors"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWT service
type JWT struct {
}

// NewJWT func
func NewJWT() *JWT {
	return &JWT{}
}

// Create func
func (j JWT) Create(data interface{}) (string, error) {
	secret := os.Getenv("APP_KEY")
	if secret == "" {
		secret = "2Yu4i1lTSrmigPyb9RYxYJ35WcnxDOQsCBCOTfoo2Yu4i1lTSrmigPyb9RYx"
	}
	token := jwt.New(jwt.SigningMethodHS256)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.New("Token generation fail")
	}

	return t, nil
}
