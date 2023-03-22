package messaging

import (
	"log"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
)

// APIMessaging struct
type APIMessaging struct {
	Config  *gowok.Config
	Service *service.Service
}

func NewAPIMessaging() *APIMessaging {
	conf := ioc.Get(gowok.Config{})
	service := ioc.Get(service.Service{})

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
