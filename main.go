package main

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/api/messaging"
	"github.com/hadihammurabi/belajar-go-rest-api/api/web"
)

func main() {
	project := gowok.Get()
	project.Runner.AddRunFunc(func() {
		web.ConfigureRoute()
		messaging.ConfigureMessage()
	})
	project.Runner.Run()
}
