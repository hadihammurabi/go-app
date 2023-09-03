package api

import (
	"github.com/hadihammurabi/belajar-go-rest-api/api/grpc"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
)

func Run() {
	go (rest.Get()).Run()
	go (grpc.Get()).Run()
	go (messaging.Get()).Run()
}
