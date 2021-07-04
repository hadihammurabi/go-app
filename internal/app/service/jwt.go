package service

import (
	"context"
	"errors"
	"time"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/app/entity"
	"github.com/hadihammurabi/belajar-go-rest-api/util"
	"github.com/sarulabs/di"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTService interface
type JWTService interface {
	Create(*entity.User) (*entity.Token, error)
	GetUser(context.Context, string) (*entity.User, error)
}

// jwtService struct
type jwtService struct {
	Config       *config.JWTConfig
	UserService  UserService
	TokenService TokenService
}

// NewJWTService func
func NewJWTService(ioc di.Container) JWTService {
	jwtConfig := ioc.Get("config").(*(config.Config)).JWT

	return &jwtService{
		Config:       jwtConfig,
		UserService:  NewUserService(ioc),
		TokenService: NewTokenService(ioc),
	}
}

// Create func
func (s jwtService) Create(userData *entity.User) (*entity.Token, error) {
	claims := &jwt.StandardClaims{}
	claims.ExpiresAt = time.Now().Add(time.Hour * 3).Unix()
	t, err := util.CreateJWTWithClaims(s.Config.Secret, claims)
	if err != nil {
		return nil, errors.New("token generation fail")
	}

	expToTime := time.Unix(claims.ExpiresAt, 0)
	tokenCreated, err := s.TokenService.Create(&entity.Token{
		UserID:    userData.ID,
		Token:     t,
		ExpiredAt: &expToTime,
	})
	if err != nil {
		return nil, errors.New("token generation fail")
	}

	return tokenCreated, nil
}

// GetUser func
func (s jwtService) GetUser(c context.Context, token string) (*entity.User, error) {
	tokenData, err := s.TokenService.FindByToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.UserService.FindByID(c, tokenData.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
