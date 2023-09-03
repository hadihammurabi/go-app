package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/entity"
)

// JWTService interface
type JWTService interface {
	Create(*entity.User) (*entity.Token, error)
	GetUser(context.Context, string) (*entity.User, error)
}

// jwtService struct
type jwtService struct {
	Config       *gowok.Config
	UserService  UserService
	TokenService TokenService
	// Cache        *cache.Redis
}

// NewJWTService func
func NewJWTService(config *gowok.Config, repo *repository.Repository) JWTService {
	return jwtService{
		Config:       config,
		UserService:  NewUserService(config, repo),
		TokenService: NewTokenService(repo),
		// Cache:        config.Redis,
	}
}

// Create func
func (s jwtService) Create(userData *entity.User) (*entity.Token, error) {
	claims := &entity.JWTClaims{
		UserID: userData.ID,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	jwk := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := jwk.SignedString([]byte(s.Config.Security.Secret))
	if err != nil {
		return nil, errors.New("token generation fail")
	}

	// if s.Cache != nil {
	// 	userData.CreatedAt = nil
	// 	s.Cache.Set(stringUtil.ToCacheKey("auth", "token", t), userData, 3*time.Hour)
	// }

	return &entity.Token{
		Token: t,
	}, nil
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
