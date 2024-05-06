package web

import (
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/exception"
)

func Configure(project *gowok.Project) {
	web := project.Web
	redis := project.Cache().OrPanic(exception.ErrNoDatabaseFound)

	index := web.Group("/")
	index.Get("", Index(redis))
}
