package grpc

import (
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc/index"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/api"
)

var a *api.Grpc

func ConfigureServices(api *api.Grpc) {
	index.RegisterIndexServer(api.Grpc, index.New())
}

func Get() *api.Grpc {
	if a != nil {
		return a
	}

	a = api.NewAPIGrpc()
	ConfigureServices(a)
	return a
}
