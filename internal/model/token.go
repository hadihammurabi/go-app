package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Token model
type Token struct {
	Base
	UserID    uuid.UUID  `json:"user_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// BeforeCreate func
func (u *Token) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}

// JWTClaims model
type JWTClaims struct {
	*jwt.StandardClaims
	UserID uuid.UUID `json:"user_id,omitempty"`
}
