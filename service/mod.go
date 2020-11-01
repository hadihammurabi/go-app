package service

import "gorm.io/gorm"

// Service struct
type Service struct {
	Auth *Auth
	User *User
	JWT  *JWT
}

// NewService func
func NewService(database *gorm.DB) (service *Service) {
	service = &Service{
		Auth: NewAuth(database),
		User: NewUser(database),
		JWT:  NewJWT(database),
	}
	return
}
