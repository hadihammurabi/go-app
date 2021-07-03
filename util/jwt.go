package util

import "github.com/dgrijalva/jwt-go"

func CreateJWTWithClaims(secret string, claims *jwt.StandardClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
