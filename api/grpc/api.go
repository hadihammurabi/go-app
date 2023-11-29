package grpc

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc/index"
)

func ConfigureServices() {
	GRPC := gowok.Get().GRPC
	index.RegisterIndexServer(GRPC, index.New())
}
