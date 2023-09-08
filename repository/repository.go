package repository

import (
	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
)

// Repository struct
type Repository struct {
}

var r *Repository

// NewRepository func
func NewRepository(db *gowok.SQL) *Repository {
	return &Repository{}
}

func Get() *Repository {
	if r != nil {
		return r
	}

	dr := driver.Get()
	r = NewRepository(dr.SQL)
	return r
}
