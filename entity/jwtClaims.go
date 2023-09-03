package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// JWTClaims model
type JWTClaims struct {
	*jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id,omitempty"`
}
