package grpc

import (
	"log"
	"net"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc/index"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/internal/service"
	"github.com/hadihammurabi/go-ioc/ioc"
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
	internalApp := ioc.Get(internal.App{})
	service := internalApp.Service

	api := &APIGrpc{
		Config:  conf,
		Service: service,
		Grpc:    grpc.NewServer(),
	}
	return api
}

func (d *APIGrpc) Run() {
	index.RegisterIndexServer(d.Grpc, index.New())

	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		panic(err)
	}

	log.Println("API GRPC started")
	err = d.Grpc.Serve(listen)
	if err != nil {
		panic(err)
	}
}

func (d *APIGrpc) Stop() {
	d.Grpc.GracefulStop()
	log.Println("GRPC was stopped")
}
