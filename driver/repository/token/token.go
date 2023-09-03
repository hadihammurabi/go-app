package token

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"gorm.io/gorm"
)

// TokenRepository interface
type TokenRepository interface {
	Create(*entity.Token) (*entity.Token, error)
	FindByUserID(uuid.UUID) (*entity.Token, error)
	FindByToken(string) (*entity.Token, error)
}

// New func
func New(db *gorm.DB) TokenRepository {
	return newSQL(db)
}
