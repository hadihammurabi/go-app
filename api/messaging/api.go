package messaging

import (
	"log"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
)

// APIMessaging struct
type APIMessaging struct {
	Config  *gowok.Config
	Service *service.Service
}

var a *APIMessaging

func NewAPIMessaging() *APIMessaging {
	conf := driver.Get().Config
	sv := service.Get()

	api := &APIMessaging{
		Config:  conf,
		Service: sv,
	}
	return api
}

func Get() *APIMessaging {
	if a != nil {
		return a
	}

	a = NewAPIMessaging()
	return a
}

func (d *APIMessaging) Run() {
	log.Println("API messaging started")
	// d.Hello()
}

func (d *APIMessaging) Stop() {
	log.Println("Messaging was stopped")
}
