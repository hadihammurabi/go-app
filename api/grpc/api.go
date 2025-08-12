package grpc

import (
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc/index"
)

func Configure(project *gowok.Project) {
	index.RegisterIndexServer(grpc.Server(), index.New())
}
