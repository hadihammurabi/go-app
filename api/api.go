package api

import (
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
)

func prepareAll() {
	apiRest := rest.NewAPIRest()
	ioc.Set(func() rest.APIRest {
		return *apiRest
	})

	apiMessaging := messaging.NewAPIMessaging()
	ioc.Set(func() messaging.APIMessaging {
		return *apiMessaging
	})

	apiGrpc := grpc.NewAPIGrpc()
	ioc.Set(func() grpc.APIGrpc {
		return *apiGrpc
	})
}

func Run() {
	prepareAll()

	go (ioc.Get(rest.APIRest{})).Run()
	go (ioc.Get(grpc.APIGrpc{})).Run()
	go (ioc.Get(messaging.APIMessaging{})).Run()
}
