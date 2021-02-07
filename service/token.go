package service

import (
	"belajar-go-rest-api/entity"
	"belajar-go-rest-api/repository"

	uuid "github.com/satori/go.uuid"
)

// TokenService struct
type TokenService struct {
	repo *repository.Repository
}

// NewTokenService func
func NewTokenService(repo *repository.Repository) entity.TokenService {
	return &TokenService{
		repo: repo,
	}
}

// Create func
func (u TokenService) Create(token *entity.Token) (*entity.Token, error) {
	return u.repo.Token.Create(token)
}

// FindByUserID func
func (u TokenService) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	return u.repo.Token.FindByUserID(id)
}
