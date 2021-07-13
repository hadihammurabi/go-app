package repository

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
	"github.com/sarulabs/di"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// TokenRepository interface
type TokenRepository interface {
	Create(*model.Token) (*model.Token, error)
	FindByUserID(uuid.UUID) (*model.Token, error)
	FindByToken(string) (*model.Token, error)
}

// tokenRepository struct
type tokenRepository struct {
	db *gorm.DB
}

// NewTokenRepository func
func NewTokenRepository(ioc di.Container) TokenRepository {
	database := getDatabase(ioc)

	return &tokenRepository{
		db: database,
	}
}

// Create func
func (u tokenRepository) Create(token *model.Token) (*model.Token, error) {
	err := u.db.Create(&token).Error
	return token, err
}

// FindByUserID func
func (u tokenRepository) FindByUserID(id uuid.UUID) (*model.Token, error) {
	token := &model.Token{}
	err := u.db.Where("id = ?", id).First(&token).Error
	return token, err
}

// FindByToken func
func (u tokenRepository) FindByToken(token string) (*model.Token, error) {
	tokenDB := &model.Token{}
	err := u.db.Where("token = ?", token).First(&tokenDB).Error
	return tokenDB, err
}
