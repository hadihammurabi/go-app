package messaging

import (
	"log"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/go-ioc/ioc"
)

// APIMessaging struct
type APIMessaging struct {
	Config  *config.Config
	Service *service.Service
}

func NewAPIMessaging() *APIMessaging {
	conf := ioc.Get(config.Config{})
	internalApp := ioc.Get(internal.App{})
	service := internalApp.Service

	api := &APIMessaging{
		Config:  conf,
		Service: service,
	}
	return api
}

func (d *APIMessaging) Run() {
	// d.Hello()
}

func (d *APIMessaging) Stop() {
	log.Println("Messaging was stopped")
}
