package runner

import (
	"os"
	"os/signal"
	"syscall"
)

func GracefulStop(gracefulStop chan os.Signal, callback func()) {
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go callback()
}
