package delivery

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/mq"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

type Delivery struct {
	Rest *rest.Delivery
	MQ   *mq.Delivery
}

func NewDelivery(ioc di.IOC) *Delivery {
	deliveryRest := rest.Init(ioc)
	deliveryMQ := mq.Init(ioc)

	return &Delivery{
		Rest: deliveryRest,
		MQ:   deliveryMQ,
	}
}

func (d Delivery) Run() {
	go d.Rest.Run()
	if d.MQ != nil {
		go d.MQ.Run()
	}
}
