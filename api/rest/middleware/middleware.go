package middleware

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/go-ioc/ioc"
)

// Middlewares type
type Middlewares struct {
	config  *config.Config
	service *service.Service
}

func NewMiddleware() Middlewares {
	config := ioc.Get(config.Config{})
	service := ioc.Get(service.Service{})

	return Middlewares{
		config:  config,
		service: service,
	}
}
