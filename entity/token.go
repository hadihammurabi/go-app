package entity

import (
	"time"

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

// TokenRepository interface
type TokenRepository interface {
	Create(*Token) (*Token, error)
	FindByUserID(uuid.UUID) (*Token, error)
}

// TokenService interface
type TokenService interface {
	Create(*Token) (*Token, error)
	FindByUserID(uuid.UUID) (*Token, error)
}

// CreateTokenDTO model
type CreateTokenDTO struct {
	UserID    uuid.UUID `json:"user_id,omitempty"`
	Token     string    `json:"token,omitempty"`
	ExpiredAt int       `json:"expired_at,omitempty"`
}
