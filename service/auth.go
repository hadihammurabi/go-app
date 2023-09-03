package service

import (
	"context"
	"errors"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/hash"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"gorm.io/gorm"
)

type AuthService struct {
	userService  UserService
	tokenService TokenService
	jwtService   JwtService
	config       *gowok.Config
}

// NewAuthService func
func NewAuthService(config *gowok.Config, db *gorm.DB, repo *repository.Repository) AuthService {
	return AuthService{
		userService:  NewUserService(config, db, repo),
		tokenService: NewTokenService(db, repo),
		jwtService:   NewJWTService(config, db, repo),
		config:       config,
	}
}

// Login func
func (a AuthService) Login(c context.Context, userInput *entity.User) (string, error) {
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
