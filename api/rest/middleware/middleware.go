package middleware

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
)

// Middlewares type
type Middlewares struct {
	config  *gowok.Config
	service *service.Service
}

func NewMiddleware() Middlewares {
	config := driver.Get().Config
	sv := service.Get()

	return Middlewares{
		config:  config,
		service: sv,
	}
}
