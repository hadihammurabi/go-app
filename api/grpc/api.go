package grpc

import (
	"log"
	"net"

	"github.com/gowok/gowok"
	"github.com/gowok/ioc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc/index"
	"github.com/hadihammurabi/belajar-go-rest-api/service"
	"google.golang.org/grpc"
)

// APIGrpc struct
type APIGrpc struct {
	Config  *gowok.Config
	Service *service.Service

	Grpc *grpc.Server
}

func NewAPIGrpc() *APIGrpc {
	conf := ioc.Get(gowok.Config{})
	service := ioc.Get(service.Service{})

	api := &APIGrpc{
		Config:  conf,
		Service: service,
		Grpc:    grpc.NewServer(),
	}
	return api
}

func (d *APIGrpc) Run() {
	if !d.Config.App.Grpc.Enabled {
		return
	}

	index.RegisterIndexServer(d.Grpc, index.New())

	listen, err := net.Listen("tcp", d.Config.App.Grpc.Host)
	if err != nil {
		panic(err)
	}

	log.Println("API GRPC started at", d.Config.App.Grpc.Host)
	err = d.Grpc.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (d *APIGrpc) Stop() {
	d.Grpc.GracefulStop()
	log.Println("GRPC was stopped")
}
