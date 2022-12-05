package service

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/hadihammurabi/belajar-go-rest-api/config"
// 	"github.com/hadihammurabi/belajar-go-rest-api/internal/entity"
// 	jwtUtil "github.com/hadihammurabi/belajar-go-rest-api/pkg/util/jwt"
// 	"github.com/gowok/ioc"
// )

// // JWTService interface
// type JWTService interface {
// 	Create(*entity.User) (*entity.Token, error)
// 	GetUser(context.Context, string) (*entity.User, error)
// }

// // jwtService struct
// type jwtService struct {
// 	JWTConfig config.JWTConfig
// 	// Cache        *cache.Redis
// 	UserService  UserService
// 	TokenService TokenService
// }

// // NewJWTService func
// func NewJWTService() JWTService {
// 	config := ioc.Get(config.Config{})

// 	return jwtService{
// 		JWTConfig: config.JWT,
// 		// Cache:        config.Redis,
// 		UserService:  NewUserService(),
// 		TokenService: NewTokenService(),
// 	}
// }

// // Create func
// func (s jwtService) Create(userData *entity.User) (*entity.Token, error) {
// 	claims := &entity.JWTClaims{
// 		UserID: userData.ID,
// 		StandardClaims: &jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
// 		},
// 	}
// 	t, err := jwtUtil.CreateJWTWithClaims(s.JWTConfig.Secret, claims)
// 	if err != nil {
// 		return nil, errors.New("token generation fail")
// 	}

// 	// if s.Cache != nil {
// 	// 	userData.CreatedAt = nil
// 	// 	s.Cache.Set(stringUtil.ToCacheKey("auth", "token", t), userData, 3*time.Hour)
// 	// }

// 	return &entity.Token{
// 		Token: t,
// 	}, nil
// }

// // GetUser func
// func (s jwtService) GetUser(c context.Context, token string) (*entity.User, error) {
// 	tokenData, err := s.TokenService.FindByToken(token)
// 	if err != nil {
// 		return nil, err
// 	}

// 	user, err := s.UserService.FindByID(c, tokenData.UserID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
