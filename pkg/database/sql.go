package database

import (
	"errors"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
)

func configureSql(conf config.Database) error {
	var err error
	if conf.Driver == "postgresql" {
		var db *database.PostgreSQL
		db, err = database.NewPostgresql(conf)
		ioc.Set(func() database.PostgreSQL { return *db })
	} else if conf.Driver == "mysql" {
		var db *database.MySQL
		db, err = database.NewMysql(conf)
		ioc.Set(func() database.MySQL { return *db })
	} else if conf.Driver == "sqlite" {
		var db *database.SQLite
		db, err = database.NewSqlite(conf)
		ioc.Set(func() database.SQLite { return *db })
	} else {
		err = errors.New("unknown database driver")
	}

	return err
}
