package validator

import (
	"github.com/gowok/gowok"
)

func Configure() gowok.Validator {
	v := gowok.NewValidator()
	return *v
}
