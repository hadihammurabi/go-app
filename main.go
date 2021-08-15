package main

import (
	"log"
	"os"
	"runtime"
	"syscall"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/runner"

	"net/http"
	_ "net/http/pprof"

	"github.com/joho/godotenv"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}

	_ = godotenv.Load()

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ioc := internal.NewIOC(conf)
	app := ioc[di.DI_APP].(*internal.App)

	var gracefulStop = make(chan os.Signal)
	runner.GracefulStop(gracefulStop, func() {
		sig := <-gracefulStop
		log.Printf("Caught SIG: %+v\n", sig)
		log.Println("Wait to finishing process")
		app.Delivery.Rest.HTTP.Shutdown()
		os.Exit(0)
	})

	forever := make(chan bool)
	go http.ListenAndServe("localhost:6060", nil)
	go app.Delivery.Run()
	<-forever
}
