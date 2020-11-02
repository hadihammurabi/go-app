package entities

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims struct
type JWTClaims struct {
	jwt.StandardClaims
	User *User `json:"user"`
}

// JWTService interface
type JWTService interface {
	Create(userData *User) (string, error)
}
