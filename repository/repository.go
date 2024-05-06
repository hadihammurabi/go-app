package repository

import "github.com/gowok/gowok"

// Repository struct
type Repository struct {
}

// NewRepository func
func NewRepository() Repository {
	return Repository{}
}

var Get = gowok.Singleton(NewRepository)
