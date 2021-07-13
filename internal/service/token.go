package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/repository"
	"github.com/sarulabs/di"

	uuid "github.com/satori/go.uuid"
)

// TokenService interface
type TokenService interface {
	Create(*entity.Token) (*entity.Token, error)
	FindByUserID(uuid.UUID) (*entity.Token, error)
	FindByToken(string) (*entity.Token, error)
}

// tokenService struct
type tokenService struct {
	repo *repository.Repository
}

// NewTokenService func
func NewTokenService(ioc di.Container) TokenService {
	repo := getRepository(ioc)
	return &tokenService{
		repo,
	}
}

// Create func
func (u tokenService) Create(token *entity.Token) (*entity.Token, error) {
	return u.repo.Token.Create(token)
}

// FindByUserID func
func (u tokenService) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	return u.repo.Token.FindByUserID(id)
}

// FindByToken func
func (u tokenService) FindByToken(token string) (*entity.Token, error) {
	return u.repo.Token.FindByToken(token)
}
