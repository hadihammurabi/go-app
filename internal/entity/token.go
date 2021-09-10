package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/db/table"
)

// Token entity
type Token struct {
	UserID    uuid.UUID  `json:"user_id,omitempty"`
	Token     string     `json:"token,omitempty"`
	ExpiredAt *time.Time `json:"expired_at,omitempty"`
}

// ToTable func
func (u Token) ToTable() *table.Token {
	return &table.Token{
		UserID:    u.UserID,
		Token:     u.Token,
		ExpiredAt: u.ExpiredAt,
	}
}

// TokenFromTable func
func TokenFromTable(fromTable *table.Token) *Token {
	return &Token{
		UserID:    fromTable.UserID,
		Token:     fromTable.Token,
		ExpiredAt: fromTable.ExpiredAt,
	}
}
