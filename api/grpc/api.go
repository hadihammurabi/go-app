package grpc

import "github.com/hadihammurabi/belajar-go-rest-api/driver/api"

var a *api.APIGrpc

func Get() *api.APIGrpc {
	if a != nil {
		return a
	}

	a = api.NewAPIGrpc()
	return a
}
