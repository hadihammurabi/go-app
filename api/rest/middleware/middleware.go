package middleware

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/go-ioc/ioc"
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
