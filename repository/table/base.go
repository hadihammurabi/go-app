package table

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key;" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}
