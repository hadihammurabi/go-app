package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository/token"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository/user"
	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	User  user.UserRepository
	Token token.TokenRepository
}

// NewRepository func
func NewRepository(db *gorm.DB) Repository {
	return Repository{
		User:  user.New(db),
		Token: token.New(db),
	}
}
