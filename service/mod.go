package service

import (
	"belajar-go-rest-api/repository"
)

// Service struct
type Service struct {
	Auth *Auth
	User *User
	JWT  *JWT
}

// NewService func
func NewService(repo *repository.Repository) (service *Service) {
	service = &Service{
		Auth: NewAuth(repo),
		User: NewUser(repo),
		JWT:  NewJWT(repo),
	}
	return
}
