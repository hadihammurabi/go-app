package main

import (
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
)

// NewIOC func
func NewIOC(conf *config.Config) di.IOC {
	ioc := di.IOC{}
	ioc[di.DI_CONFIG] = conf
	ioc[di.DI_APP] = internal.NewApp(ioc)
	return ioc
}
