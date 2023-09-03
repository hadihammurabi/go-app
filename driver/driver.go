package driver

import (
	"os"

	"github.com/gowok/gowok"
	"github.com/gowok/gowok/exception"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/util/runner"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
)

type Driver struct {
	Config     *gowok.Config
	SQL        *gowok.SQL
	Validator  *gowok.Validator
	Repository *repository.Repository
}

var d *Driver

func Get() *Driver {
	if d != nil {
		return d
	}

	runner.PrepareRuntime()

	conf := gowok.Must(
		gowok.NewConfig(os.OpenFile("config.yaml", os.O_RDONLY, 600)),
	)
	sqlDB := gowok.Must(
		gowok.NewSQL(conf.Databases),
	)
	val := gowok.NewValidator()

	repo := repository.NewRepository(
		sqlDB.Get().OrPanic(exception.ErrNoDatabaseFound),
	)

	d = &Driver{
		Config:     conf,
		SQL:        &sqlDB,
		Validator:  val,
		Repository: &repo,
	}
	return d
}
