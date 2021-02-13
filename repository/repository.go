package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"gorm.io/gorm"

	"github.com/sarulabs/di"
)

// Repository struct
type Repository struct {
	User  entity.UserRepository
	Token entity.TokenRepository
}

// NewRepository func
func NewRepository(ioc di.Container) (repo *Repository) {
	database := ioc.Get("database").(*gorm.DB)

	repo = &Repository{
		User:  NewUserRepository(database),
		Token: NewTokenRepository(database),
	}
	return
}
