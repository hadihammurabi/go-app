package web

import (
	"github.com/gowok/gowok"
)

func ConfigureRoute() {
	api := gowok.Get().Web
	api.Mount("", Index())
	// api.HTTP.Mount("/auth", Auth(api).router)
}
