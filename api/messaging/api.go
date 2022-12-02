package messaging

import (
	"log"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/go-ioc/ioc"
)

// APIMessaging struct
type APIMessaging struct {
	Config *config.Config
}

func NewAPIMessaging() *APIMessaging {
	conf := ioc.Get(config.Config{})

	api := &APIMessaging{
		Config: conf,
	}
	return api
}

func (d *APIMessaging) Run() {
	d.Hello()
}

func (d *APIMessaging) Stop() {
	log.Println("Messaging was stopped")
}
