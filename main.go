package main

import (
	"os"
	"runtime"
	"syscall"

	"github.com/hadihammurabi/belajar-go-rest-api/api/rest"
	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/db/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/util/runner"
	"github.com/hadihammurabi/go-ioc/ioc"
	"github.com/spf13/viper"

	"net/http"
	_ "net/http/pprof"
)

func init() {
	configLocation := os.Getenv("APP_CONFIG")
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configLocation)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

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

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ioc.Set(func() config.Config {
		return conf
	})

	repo := repository.NewRepository()
	ioc.Set(func() repository.Repository {
		return repo
	})

	app := internal.NewApp()
	ioc.Set(func() internal.App {
		return *app
	})

	apiRest := rest.NewAPIRest()
	ioc.Set(func() rest.APIRest {
		return *apiRest
	})

	forever := make(chan bool)
	var gracefulStop = make(chan os.Signal)
	runner.GracefulStop(gracefulStop, func() {
		<-gracefulStop

		forever <- true
		apiRest.Stop()
		os.Exit(0)
	})

	go http.ListenAndServe("localhost:6060", nil)
	go apiRest.Run()
	<-forever
}
