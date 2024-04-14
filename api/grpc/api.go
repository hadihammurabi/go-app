package grpc

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc/index"
)

func Configure(project *gowok.Project) {
	g := project.GRPC

	index.RegisterIndexServer(g, index.New())
}
