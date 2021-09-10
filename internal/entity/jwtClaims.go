package entity

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// JWTClaims model
type JWTClaims struct {
	*jwt.StandardClaims
	UserID uuid.UUID `json:"user_id,omitempty"`
}
