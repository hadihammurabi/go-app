package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/entity"

	"github.com/sarulabs/di"
)

// Repository struct
type Repository struct {
	User  entity.UserRepository
	Token entity.TokenRepository
}

// NewRepository func
func NewRepository(ioc di.Container) (repo *Repository) {
	repo = &Repository{
		User:  NewUserRepository(ioc),
		Token: NewTokenRepository(ioc),
	}
	return
}
