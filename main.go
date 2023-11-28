package main

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/api"
)

func main() {
	project := gowok.Get()
	web.ConfigureRoute(api.NewAPIRest())
	project.Runner.Run()
}
