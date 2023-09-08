package driver

import (
	"os"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/driver/util"
)

type Driver struct {
	Config    *gowok.Config
	SQL       *gowok.SQL
	MongoDB   *gowok.MongoDB
	Validator *gowok.Validator
}

var d *Driver

func Get() *Driver {
	if d != nil {
		return d
	}

	util.PrepareRuntime()

	conf := gowok.Must(
		gowok.NewConfig(os.OpenFile("config.yaml", os.O_RDONLY, 600)),
	)
	sqlDB := gowok.Must(
		gowok.NewSQL(conf.Databases),
	)
	mongoDB := gowok.Must(
		gowok.NewMongoDB(conf.Databases),
	)
	val := gowok.NewValidator()

	d = &Driver{
		Config:    conf,
		SQL:       &sqlDB,
		MongoDB:   &mongoDB,
		Validator: val,
	}
	return d
}
