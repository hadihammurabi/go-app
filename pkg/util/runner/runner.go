package runner

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/spf13/viper"
)

func PrepareRuntime() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rLimit syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
	rLimit.Cur = rLimit.Max
	if err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit); err != nil {
		panic(err)
	}
}

func PrepareConfig() {
	configLocation := os.Getenv("APP_CONFIG")

	if configLocation == "" {
		configLocation = "config.yaml"
	}

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configLocation)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func GracefulStop(gracefulStop chan os.Signal, callback func()) {
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go callback()
}
