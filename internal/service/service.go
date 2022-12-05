package service

// Service struct
type Service struct {
	Auth  AuthService
	User  UserService
	Token TokenService
	// JWT   JWTService
}

// NewService func
func NewService() (service *Service) {
	service = &Service{
		Auth:  NewAuthService(),
		User:  NewUserService(),
		Token: NewTokenService(),
		// JWT:   NewJWTService(),
	}
	return
}
