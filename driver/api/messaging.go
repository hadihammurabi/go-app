package api

import (
	"log"
)

// Messaging struct
type Messaging struct {
}

func NewMessaging() *Messaging {
	api := &Messaging{}
	return api
}

func (d *Messaging) Run() {
	log.Println("API messaging started")
}

func (d *Messaging) Stop() {
	log.Println("Messaging was stopped")
}
