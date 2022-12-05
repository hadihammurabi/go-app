package messaging

import (
	"log"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
)

// APIMessaging struct
type APIMessaging struct {
	Config  *gowok.Config
	Service *service.Service
}

func NewAPIMessaging() *APIMessaging {
	conf := ioc.Get(gowok.Config{})
	internalApp := ioc.Get(internal.App{})
	service := internalApp.Service

	api := &APIMessaging{
		Config:  conf,
		Service: service,
	}
	return api
}

func (d *APIMessaging) Run() {
	log.Println("API messaging started")
	// d.Hello()
}

func (d *APIMessaging) Stop() {
	log.Println("Messaging was stopped")
}
