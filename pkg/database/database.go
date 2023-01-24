package database

import (
	"fmt"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
)

func Configure(conf []config.Database) {

	for _, dbConf := range conf {
		driver, ok := MapDriver(dbConf.Driver)
		if !ok {
			panic(fmt.Sprintf("unknown database driver %s", dbConf.Driver))
		}

		if driver.Type == SQL {
			err := configureSql(dbConf)
			if err != nil {
				panic(err)
			}
		}

		if dbConf.Driver == "mongo" {
			mg, err := database.NewMongo(dbConf)
			if err != nil {
				panic(err)
			}

			ioc.Set(func() *database.Mongo { return mg })
		}

		if dbConf.Driver == "redis" {
			rdb, err := database.NewRedis(dbConf)
			if err != nil {
				panic(err)
			}

			ioc.Set(func() *database.Redis { return rdb })
		}
	}
}
