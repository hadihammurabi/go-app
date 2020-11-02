package service

import (
	"belajar-go-rest-api/entities"
	"belajar-go-rest-api/repository"
)

// Service struct
type Service struct {
	Auth entities.AuthService
	User entities.UserService
	JWT  entities.JWTService
}

// NewService func
func NewService(repo *repository.Repository) (service *Service) {
	service = &Service{
		Auth: NewAuthService(repo),
		User: NewUserService(repo),
		JWT:  NewJWTService(),
	}
	return
}
