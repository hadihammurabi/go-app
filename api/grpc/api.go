package grpc

import (
	"log"
	"net"

	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc/index"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"google.golang.org/grpc"
)

// APIGrpc struct
type APIGrpc struct {
	Grpc *grpc.Server
}

var a *APIGrpc

func NewAPIGrpc() *APIGrpc {
	api := &APIGrpc{
		Grpc: grpc.NewServer(),
	}

	index.RegisterIndexServer(api.Grpc, index.New())

	return api
}

func Get() *APIGrpc {
	if a != nil {
		return a
	}

	a = NewAPIGrpc()
	return a
}

func (d *APIGrpc) Run() {
	grpcConf := driver.Get().Config.App.Grpc
	if !grpcConf.Enabled {
		return
	}

	listen, err := net.Listen("tcp", grpcConf.Host)
	if err != nil {
		panic(err)
	}

	log.Println("API GRPC started at", grpcConf.Host)
	err = d.Grpc.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (d *APIGrpc) Stop() {
	d.Grpc.GracefulStop()
	log.Println("GRPC was stopped")
}
