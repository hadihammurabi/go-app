package token

import (
	"time"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/repository/base"
	"gorm.io/gorm"
)

// Table model
type Table struct {
	base.Table
	UserID    uuid.UUID  `json:"user_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// TableName func
func (u *Table) TableName(tx *gorm.DB) string {
	return "tokens"
}

// BeforeCreate func
func (u *Table) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewRandom()
	u.ID = id
	return
}

// ToEntity func
func (t Table) ToEntity() *entity.Token {
	return &entity.Token{
		UserID:    t.UserID,
		Token:     t.Token,
		ExpiredAt: t.ExpiredAt,
	}
}

// FromEntity func
func FromEntity(e *entity.Token) *Table {
	return &Table{
		UserID:    e.UserID,
		Token:     e.Token,
		ExpiredAt: e.ExpiredAt,
	}
}
