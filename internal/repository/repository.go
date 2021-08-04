package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	User  UserRepository
	Token TokenRepository
}

// NewRepository func
func NewRepository(ioc di.IOC) (repo *Repository) {
	repo = &Repository{
		User:  NewUserRepository(ioc),
		Token: NewTokenRepository(ioc),
	}
	return
}

func getDatabase(ioc di.IOC) *gorm.DB {
	return ioc[di.DI_CONFIG].(*config.Config).DB
}
