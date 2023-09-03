package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	All(context.Context) ([]*entity.User, error)
	Create(context.Context, *entity.User) (*entity.User, error)
	FindByID(context.Context, uuid.UUID) (*entity.User, error)
	FindByEmail(context.Context, string) (*entity.User, error)
	ChangePassword(context.Context, uuid.UUID, string) (*entity.User, error)
}

// New func
func New(db *gorm.DB) UserRepository {
	return newSQL(db)
}
