package main

import (
	"os"
	"runtime"
	"syscall"

	"github.com/hadihammurabi/belajar-go-rest-api/config"
	"github.com/hadihammurabi/belajar-go-rest-api/internal"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/di"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg/util/runner"
	"github.com/spf13/viper"

	"net/http"
	_ "net/http/pprof"

	"github.com/joho/godotenv"
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

	_ = godotenv.Load()

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	ioc := internal.NewIOC(conf)
	app := ioc[di.DI_APP].(*internal.App)

	var gracefulStop = make(chan os.Signal)
	runner.GracefulStop(gracefulStop, func() {
		<-gracefulStop
		app.APIRest.Stop()
		os.Exit(0)
	})

	forever := make(chan bool)
	go http.ListenAndServe("localhost:6060", nil)
	go app.APIRest.Run()
	<-forever
}
