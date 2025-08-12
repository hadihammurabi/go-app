package web

import (
	"github.com/gowok/gowok"
)

func Configure(project *gowok.Project) {
	web := gowok.Router()

	index := web.Group("/")
	index.Get("", Index())
}
