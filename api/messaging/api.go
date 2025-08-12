package messaging

import (
	"log/slog"

	"github.com/gowok/gowok"
)

func Configure(project *gowok.Project) {
	slog.Info("starting messaging")
	go Hello()
}
