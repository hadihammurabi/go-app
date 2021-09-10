package repository

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"

	"gorm.io/gorm"
)

// TokenRepository interface
type TokenRepository interface {
	Create(*entity.Token) (*entity.Token, error)
	FindByUserID(uuid.UUID) (*entity.Token, error)
	FindByToken(string) (*entity.Token, error)
}

// tokenRepository struct
type tokenRepository struct {
	db *gorm.DB
}

// NewTokenRepository func
func NewTokenRepository(ioc di.IOC) TokenRepository {
	database := getDatabase(ioc)

	return &tokenRepository{
		db: database,
	}
}

// Create func
func (u tokenRepository) Create(token *entity.Token) (*entity.Token, error) {
	err := u.db.Create(&token).Error
	return token, err
}

// FindByUserID func
func (u tokenRepository) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	token := &entity.Token{}
	err := u.db.Where("id = ?", id).First(&token).Error
	return token, err
}

// FindByToken func
func (u tokenRepository) FindByToken(token string) (*entity.Token, error) {
	tokenDB := &entity.Token{}
	err := u.db.Where("token = ?", token).First(&tokenDB).Error
	return tokenDB, err
}
