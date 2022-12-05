package api

import (
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
	"github.com/hadihammurabi/go-ioc/ioc"
)

func PrepareAll() {
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
