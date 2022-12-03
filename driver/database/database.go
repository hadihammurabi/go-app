package database

import (
	"errors"

	"github.com/hadihammurabi/belajar-go-rest-api/driver/database/sql"
	"gorm.io/gorm"
)

type Driver string

const (
	DriverPostgresql Driver = "postgresql"
	DriverSqlite     Driver = "sqlite"
	DriverMysql      Driver = "mysql"
)

type Config struct {
	Driver Driver
	DSN    string
}

type Database struct {
	*gorm.DB
}

func NewDatabase(config Config) (Database, error) {
	var db *gorm.DB
	var err error
	if config.Driver == DriverPostgresql {
		db, err = sql.ConfigurePostgresql(sql.Config{
			DSN: config.DSN,
		})
	} else if config.Driver == DriverMysql {
		db, err = sql.ConfigureMysql(sql.Config{
			DSN: config.DSN,
		})
	} else if config.Driver == DriverSqlite {
		db, err = sql.ConfigureSqlite(sql.Config{
			DSN: config.DSN,
		})
	} else {
		err = errors.New("unknown database driver")
	}

	if err != nil {
		return Database{}, err
	}

	return Database{
		DB: db,
	}, nil
}
