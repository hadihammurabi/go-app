package mq

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
	gorabbitmq "github.com/hadihammurabi/go-rabbitmq"
)

// Delivery struct
type Delivery struct {
	Service *service.Service
	Config  *config.Config
	MQ      gorabbitmq.MQ
}

// Init func
func Init(ioc di.IOC) (*Delivery, error) {
	service := ioc["service"].(*service.Service)
	conf := ioc["config"].(*config.Config)

	mq, err := gorabbitmq.NewMQ(conf.MQ.GetURL())
	if err != nil {
		return nil, err
	}

	delivery := &Delivery{
		Service: service,
		Config:  conf,
		MQ:      mq,
	}
	return delivery, nil
}

func (d *Delivery) Run() {
	// TODO: Call all MQ consumer here
}
