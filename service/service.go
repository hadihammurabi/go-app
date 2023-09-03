package service

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
)

// Service struct
type Service struct {
	Auth  AuthService
	User  UserService
	Token TokenService
	// JWT   JWTService
}

// NewService func
func NewService() *Service {
	dr := driver.Get()
	config := dr.Config
	repo := dr.Repository

	service := &Service{
		Auth:  NewAuthService(config, repo),
		User:  NewUserService(config, repo),
		Token: NewTokenService(repo),
		// JWT:   NewJWTService(),
	}

	return service
}

var s *Service

func Get() *Service {
	if s != nil {
		return s
	}

	s = NewService()
	return s
}
