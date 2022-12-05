package middleware

import (
	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
)

// Middlewares type
type Middlewares struct {
	config  *gowok.Config
	service *service.Service
}

func NewMiddleware() Middlewares {
	config := ioc.Get(gowok.Config{})
	service := ioc.Get(service.Service{})

	return Middlewares{
		config:  config,
		service: service,
	}
}
