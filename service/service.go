package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"github.com/sarulabs/di"
)

// Service struct
type Service struct {
	Auth  entity.AuthService
	User  entity.UserService
	Token entity.TokenService
	JWT   entity.JWTService
}

// NewService func
func NewService(ioc di.Container) (service *Service) {
	service = &Service{
		Auth:  NewAuthService(ioc),
		User:  NewUserService(ioc),
		Token: NewTokenService(ioc),
		JWT:   NewJWTService(ioc),
	}
	return
}

func getRepository(ioc di.Container) *repository.Repository {
	return ioc.Get("repository").(*repository.Repository)
}
