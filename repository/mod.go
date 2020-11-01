package repository

import "gorm.io/gorm"

// Repository struct
type Repository struct {
	User *UserRepository
}

// NewRepository func
func NewRepository(database *gorm.DB) (repo *Repository) {
	repo = &Repository{
		User: NewUserRepository(database),
	}
	return
}
