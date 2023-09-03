package table

import (
	"time"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"gorm.io/gorm"
)

type Token struct {
	Base
	UserID    uuid.UUID  `json:"user_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// TableName func
func (u *Token) TableName(tx *gorm.DB) string {
	return "tokens"
}

// BeforeCreate func
func (u *Token) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewRandom()
	u.ID = id
	return
}

// ToEntity func
func (t Token) ToEntity() *entity.Token {
	return &entity.Token{
		UserID:    t.UserID,
		Token:     t.Token,
		ExpiredAt: t.ExpiredAt,
	}
}
