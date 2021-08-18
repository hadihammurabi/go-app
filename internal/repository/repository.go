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
func NewRepository(ioc di.IOC) Repository {
	return Repository{
		User:  NewUserRepository(ioc),
		Token: NewTokenRepository(ioc),
	}
}

func getDatabase(ioc di.IOC) *gorm.DB {
	config := ioc[di.DI_CONFIG].(config.Config)
	return &config.DB
}
