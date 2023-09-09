package messaging

import "github.com/hadihammurabi/belajar-go-rest-api/driver/api"

func ConfigureMessage(api *api.Messaging) {
	go Hello()
}

var a *api.Messaging

func Get() *api.Messaging {
	if a != nil {
		return a
	}

	a = api.NewMessaging()
	ConfigureMessage(a)
	return a
}
