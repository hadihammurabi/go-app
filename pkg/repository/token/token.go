package token

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
	"github.com/hadihammurabi/go-ioc/ioc"
)

// TokenRepository interface
type TokenRepository interface {
	Create(*Table) (*Table, error)
	FindByUserID(uuid.UUID) (*Table, error)
	FindByToken(string) (*Table, error)
}

// New func
func New() TokenRepository {
	db := ioc.Get(database.Database{}).DB

	return newSQL(db)
}
