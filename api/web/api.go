package web

import (
	"github.com/gowok/gowok"
)

func Configure(project *gowok.Project) {
	web := project.Web

	web.Mount("", Index())
	// api.HTTP.Mount("/auth", Auth(api).router)
}
