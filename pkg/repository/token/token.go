package token

import (
	"github.com/google/uuid"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
)

// TokenRepository interface
type TokenRepository interface {
	Create(*entity.Token) (*entity.Token, error)
	FindByUserID(uuid.UUID) (*entity.Token, error)
	FindByToken(string) (*entity.Token, error)
}

// New func
func New() TokenRepository {
	db := ioc.Get(database.PostgreSQL{}).DB

	return newSQL(db)
}
