package entities

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims struct
type JWTClaims struct {
	jwt.StandardClaims
	User *User `json:"user"`
}
