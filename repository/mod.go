package repository

import (
	"belajar-go-rest-api/entities"

	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
	User entities.UserRepository
}

// NewRepository func
func NewRepository(database *gorm.DB) (repo *Repository) {
	repo = &Repository{
		User: NewUserRepository(database),
	}
	return
}
