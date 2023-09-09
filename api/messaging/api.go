package messaging

import "github.com/hadihammurabi/belajar-go-rest-api/driver/api"

func ConfigureMessage(api *api.Messaging) {

}

var a *api.Messaging

func Get() *api.Messaging {
	if a != nil {
		return a
	}

	a = api.NewMessaging()
	return a
}
