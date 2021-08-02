package service

import (
	"context"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/model"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/cache"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
	jwtUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/jwt"
	stringUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/string"
)

// JWTService interface
type JWTService interface {
	Create(*model.User) (*model.Token, error)
	GetUser(context.Context, string) (*model.User, error)
}

// jwtService struct
type jwtService struct {
	JWTConfig    *config.JWTConfig
	Cache        *cache.Redis
	UserService  UserService
	TokenService TokenService
}

// NewJWTService func
func NewJWTService(ioc di.IOC) JWTService {
	config := getConfig(ioc)

	return &jwtService{
		JWTConfig:    config.JWT,
		Cache:        config.Redis,
		UserService:  NewUserService(ioc),
		TokenService: NewTokenService(ioc),
	}
}

// Create func
func (s jwtService) Create(userData *model.User) (*model.Token, error) {
	claims := &model.JWTClaims{
		UserID: userData.ID,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
	}
	t, err := jwtUtil.CreateJWTWithClaims(s.JWTConfig.Secret, claims)
	if err != nil {
		return nil, errors.New("token generation fail")
	}

	if s.Cache != nil {
		userData.CreatedAt = nil
		s.Cache.Set(stringUtil.ToCacheKey("auth", "token", t), userData, 3*time.Hour)
	}

	return &model.Token{
		Token: t,
	}, nil
}

// GetUser func
func (s jwtService) GetUser(c context.Context, token string) (*model.User, error) {
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
