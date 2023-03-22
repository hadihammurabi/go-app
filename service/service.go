package service

import "github.com/gowok/ioc"

// Service struct
type Service struct {
	Auth  AuthService
	User  UserService
	Token TokenService
	// JWT   JWTService
}

// NewService func
func NewService() Service {
	service := Service{
		Auth:  NewAuthService(),
		User:  NewUserService(),
		Token: NewTokenService(),
		// JWT:   NewJWTService(),
	}

	return service
}

func PrepareAll() {
	ioc.Set(func() Service { return NewService() })
}
