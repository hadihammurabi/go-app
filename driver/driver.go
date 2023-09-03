package driver

import (
	"fmt"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/driver"
	"github.com/gowok/gowok/exception"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/config"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/repository"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/util/runner"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/validator"
)

type Driver struct {
	Config     *gowok.Config
	SQL        *driver.SQL
	Validator  *gowok.Validator
	Repository *repository.Repository
}

var d *Driver

func Get() *Driver {
	if d != nil {
		return d
	}

	runner.PrepareRuntime()

	conf := config.Configure()
	sqlDB := driver.NewSQL(conf.Databases)
	val := validator.Configure()

	dbDefault := sqlDB.Get().OrPanic(exception.ErrNoDatabaseFound)
	repo := repository.NewRepository(&dbDefault)

	fmt.Println("welcom", sqlDB, val, repo)
	d = &Driver{
		Config:     conf,
		SQL:        &sqlDB,
		Validator:  &val,
		Repository: &repo,
	}
	return d
}
