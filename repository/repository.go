package repository

import (
	"belajar-go-rest-api/entity"

	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	User  entity.UserRepository
	Token entity.TokenRepository
}

// NewRepository func
func NewRepository(database *gorm.DB) (repo *Repository) {
	repo = &Repository{
		User:  NewUserRepository(database),
		Token: NewTokenRepository(database),
	}
	return
}
