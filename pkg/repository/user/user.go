package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
	"github.com/hadihammurabi/go-ioc/ioc"
)

// UserRepository interface
type UserRepository interface {
	All(context.Context) ([]*Table, error)
	Create(context.Context, *Table) (*Table, error)
	FindByID(context.Context, uuid.UUID) (*Table, error)
	FindByEmail(context.Context, string) (*Table, error)
	ChangePassword(context.Context, uuid.UUID, string) (*Table, error)
}

// New func
func New() UserRepository {
	db := ioc.Get(database.Database{}).DB

	return newSQL(db)
}
