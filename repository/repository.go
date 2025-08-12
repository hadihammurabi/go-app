package repository

import (
	"github.com/gowok/gowok/singleton"
)

// Repository struct
type Repository struct {
}

// NewRepository func
func NewRepository() Repository {
	return Repository{}
}

var Get = singleton.New(NewRepository)
