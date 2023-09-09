package driver

import (
	"os"

	"github.com/gowok/gowok"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
)

type Driver struct {
	Config    *gowok.Config
	SQL       *gowok.SQL
	MongoDB   *gowok.MongoDB
	Redis     *gowok.Redis
	Validator *gowok.Validator
}

var d *Driver

func Get() *Driver {
	if d != nil {
		return d
	}

	pkg.PrepareRuntime()

	conf := gowok.Must(
		gowok.NewConfig(os.OpenFile("config.yaml", os.O_RDONLY, 600)),
	)
	sqlDB := gowok.Must(
		gowok.NewSQL(conf.Databases),
	)
	mongoDB := gowok.Must(
		gowok.NewMongoDB(conf.Databases),
	)
	redisDB := gowok.Must(
		gowok.NewRedis(conf.Databases),
	)
	val := gowok.NewValidator()

	d = &Driver{
		Config:    conf,
		SQL:       &sqlDB,
		MongoDB:   &mongoDB,
		Redis:     &redisDB,
		Validator: val,
	}
	return d
}
