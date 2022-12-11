package database

import (
	"fmt"

	"github.com/gowok/gowok/config"
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
			err := configureMongo(dbConf)
			if err != nil {
				panic(err)
			}
		}

		if dbConf.Driver == "redis" {
			err := configureRedis(dbConf)
			if err != nil {
				panic(err)
			}
		}
	}
}
