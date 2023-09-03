package repository

import (
	"gorm.io/gorm"
)

// Repository struct
type Repository struct {
}

// NewRepository func
func NewRepository(db *gorm.DB) Repository {
	return Repository{}
}
