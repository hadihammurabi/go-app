package jwt

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
)

func CreateJWTWithClaims(secret string, claims *model.JWTClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func JWTFromHeader(auth string) (string, string, error) {
	authSplitted := strings.Split(auth, " ")
	if len(authSplitted) == 2 {
		return authSplitted[0], authSplitted[1], nil
	}

	return "", "", errors.New("invalid token")
}

func VerifyJWT(token string, secret string) (*jwt.Token, error) {
	parsed, err := jwt.ParseWithClaims(token, &model.JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func GetJWTData(token string, secret string) (*model.JWTClaims, error) {
	tokenVerified, err := VerifyJWT(token, secret)
	if err != nil {
		return nil, err
	}

	claims, ok := tokenVerified.Claims.(*model.JWTClaims)
	if !ok && !tokenVerified.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
