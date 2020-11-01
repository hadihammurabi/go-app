package repository

import "gorm.io/gorm"

// Repository struct
type Repository struct {
	User *User
}

// NewRepository func
func NewRepository(database *gorm.DB) (repo *Repository) {
	repo = &Repository{
		User: NewUser(database),
	}
	return
}
