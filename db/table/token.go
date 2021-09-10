package table

import (
	"time"

	"github.com/google/uuid"
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
	id, err := uuid.NewRandom()
	u.ID = id
	return
}
