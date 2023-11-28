package web

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/api"
)

func ConfigureRoute(api *api.Rest) {
	api.HTTP.Mount("", Index(api))
	// api.HTTP.Mount("/auth", Auth(api).router)
}
