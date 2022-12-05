package token

import (
	"github.com/google/uuid"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
)

// TokenRepository interface
type TokenRepository interface {
	Create(*entity.Token) (*entity.Token, error)
	FindByUserID(uuid.UUID) (*entity.Token, error)
	FindByToken(string) (*entity.Token, error)
}

// New func
func New() TokenRepository {
	db := ioc.Get(database.Database{}).DB

	return newSQL(db)
}
