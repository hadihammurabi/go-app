package service

import (
	"github.com/gowok/gowok/exception"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
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
	sql := dr.SQL.Get().OrPanic(exception.ErrNoDatabaseFound)
	repo := repository.Get()

	service := &Service{
		Auth:  NewAuthService(config, sql, repo),
		User:  NewUserService(config, sql, repo),
		Token: NewTokenService(sql, repo),
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
