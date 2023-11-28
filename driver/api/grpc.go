package api

import (
	"github.com/gowok/gowok"
	"google.golang.org/grpc"
)

// Grpc struct
type Grpc struct {
	Grpc *grpc.Server
}

func NewAPIGrpc() *Grpc {
	api := &Grpc{
		Grpc: gowok.Get().GRPC,
	}

	return api
}
