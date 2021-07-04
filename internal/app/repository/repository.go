package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"gorm.io/gorm"

	"github.com/sarulabs/di"
)

// Repository struct
type Repository struct {
	User  UserRepository
	Token TokenRepository
}

// NewRepository func
func NewRepository(ioc di.Container) (repo *Repository) {
	repo = &Repository{
		User:  NewUserRepository(ioc),
		Token: NewTokenRepository(ioc),
	}
	return
}

func getDatabase(ioc di.Container) *gorm.DB {
	return ioc.Get("config").(*config.Config).DB
}
