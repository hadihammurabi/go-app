package delivery

import (
	"github.com/hadihammurabi/belajar-go-rest-api/internal/delivery/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

type Delivery struct {
	Rest *rest.Delivery
}

func NewDelivery(ioc di.IOC) *Delivery {
	deliveryRest := rest.Init(ioc)

	return &Delivery{
		Rest: deliveryRest,
	}
}

func (d Delivery) Run() {
	go d.Rest.Run()
}
