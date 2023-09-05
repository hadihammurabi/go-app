package rest

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/api"
)

func ConfigureRoute(api *api.Rest) {
	api.HTTP.Mount("", Index(api))
	// api.HTTP.Mount("/auth", Auth(api).router)
}

var a *api.Rest

func Get() *api.Rest {
	if a != nil {
		return a
	}

	a = api.NewAPIRest()
	ConfigureRoute(a)
	return a
}
