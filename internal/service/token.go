package service

import (
	"github.com/google/uuid"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

// TokenService interface
type TokenService interface {
	Create(*model.Token) (*model.Token, error)
	FindByUserID(uuid.UUID) (*model.Token, error)
	FindByToken(string) (*model.Token, error)
}

// tokenService struct
type tokenService struct {
	repo *repository.Repository
}

// NewTokenService func
func NewTokenService(ioc di.IOC) TokenService {
	repo := getRepository(ioc)
	return &tokenService{
		repo,
	}
}

// Create func
func (u tokenService) Create(token *model.Token) (*model.Token, error) {
	return u.repo.Token.Create(token)
}

// FindByUserID func
func (u tokenService) FindByUserID(id uuid.UUID) (*model.Token, error) {
	return u.repo.Token.FindByUserID(id)
}

// FindByToken func
func (u tokenService) FindByToken(token string) (*model.Token, error) {
	return u.repo.Token.FindByToken(token)
}
