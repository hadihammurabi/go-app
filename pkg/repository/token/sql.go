package token

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"

	"gorm.io/gorm"
)

// tokenSQL struct
type tokenSQL struct {
	db *gorm.DB
}

// newSQL func
func newSQL(db *gorm.DB) *tokenSQL {
	return &tokenSQL{
		db: db,
	}
}

// Create func
func (u tokenSQL) Create(token *entity.Token) (*entity.Token, error) {
	err := u.db.Create(FromEntity(token)).Error
	return token, err
}

// FindByUserID func
func (u tokenSQL) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	row := &Table{}
	err := u.db.Where("id = ?", id).First(&row).Error
	return row.ToEntity(), err
}

// FindByToken func
func (u tokenSQL) FindByToken(token string) (*entity.Token, error) {
	row := &Table{}
	err := u.db.Where("token = ?", token).First(&row).Error
	return row.ToEntity(), err
}
