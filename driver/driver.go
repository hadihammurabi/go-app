package driver

import (
	"flag"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/gowok/gowok"
	"github.com/gowok/gowok/optional"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type getterByName[T any] func(name ...string) optional.Optional[T]

type Driver struct {
	Config    *gowok.Config
	SQL       getterByName[*gorm.DB]
	MongoDB   getterByName[*mongo.Client]
	Redis     getterByName[*redis.Client]
	Validator *gowok.Validator
}

var d *Driver

func Get() *Driver {
	if d != nil {
		return d
	}

	pkg.PrepareRuntime()
	configPath := flag.String("config", "config.yaml", "")
	flag.Parse()

	conf := gowok.Must(
		gowok.NewConfig(os.OpenFile(*configPath, os.O_RDONLY, 600)),
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
		SQL:       sqlDB.Get,
		MongoDB:   mongoDB.Get,
		Redis:     redisDB.Get,
		Validator: val,
	}
	return d
}
