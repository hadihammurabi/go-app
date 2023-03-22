package entity

import (
	"time"

	"github.com/google/uuid"
)

// Token entity
type Token struct {
	UserID    uuid.UUID  `json:"user_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}
