package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/database"
	"github.com/hadihammurabi/go-ioc/ioc"
	"github.com/spf13/viper"
)

// ConfigureDatabase func
func ConfigureDatabase() database.Database {
	dbconfigsFromConfig := viper.Get("database").(map[string]any)
	driver, _ := dbconfigsFromConfig["driver"].(string)
	dsn, _ := dbconfigsFromConfig["dsn"].(string)

	dbconf := database.Config{
		Driver: database.Driver(driver),
		DSN:    dsn,
	}

	db, err := database.NewDatabase(dbconf)
	if err != nil {
		panic(err)
	}

	ioc.Set(func() database.Database { return db })

	return db
}
