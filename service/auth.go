package service

import (
	"context"
	"errors"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/hash"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
)

// AuthService interface
type AuthService interface {
	Login(context.Context, *entity.User) (string, error)
}

// authService struct
type authService struct {
	userService  UserService
	tokenService TokenService
	jwtService   JWTService
	config       *gowok.Config
}

// NewAuthService func
func NewAuthService(config *gowok.Config, repo *repository.Repository) AuthService {
	return authService{
		userService:  NewUserService(config, repo),
		tokenService: NewTokenService(repo),
		jwtService:   NewJWTService(config, repo),
		config:       config,
	}
}

// Login func
func (a authService) Login(c context.Context, userInput *entity.User) (string, error) {
	user, err := a.userService.FindByEmail(c, userInput.Email)
	if err != nil {
		return "", errors.New("email or password invalid")
	}

	isPasswordValid := hash.PasswordVerify(userInput.Password, user.Password, a.config.App.Key)
	if isPasswordValid {
		return "", errors.New("email or password invalid")
	}

	token, err := a.jwtService.Create(user)
	if err != nil {
		return "", errors.New("email or password invalid")
	}

	return token.Token, nil
}
