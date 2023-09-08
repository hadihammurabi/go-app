package api

import (
	"log"
	"net"

	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"google.golang.org/grpc"
)

// Grpc struct
type Grpc struct {
	Grpc *grpc.Server
}

func NewAPIGrpc() *Grpc {
	api := &Grpc{
		Grpc: grpc.NewServer(),
	}

	return api
}

func (d *Grpc) Run() {
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

func (d *Grpc) Stop() {
	d.Grpc.GracefulStop()
	log.Println("GRPC was stopped")
}
