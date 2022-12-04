package database

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
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
	db, err := config.configure()
	if err != nil {
		return Database{}, err
	}

	return Database{
		DB: db,
	}, nil
}

func (c Config) configure() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	if c.Driver == DriverPostgresql {
		db, err = gorm.Open(postgres.Open(c.DSN), &gorm.Config{})
	} else if c.Driver == DriverMysql {
		db, err = gorm.Open(mysql.Open(c.DSN), &gorm.Config{})
	} else if c.Driver == DriverSqlite {
		location := c.DSN
		if location == "" {
			location = "db.sqlite3"
		}

		db, err = gorm.Open(sqlite.Open(location), &gorm.Config{})
	} else {
		err = errors.New("unknown database driver")
	}

	return db, err
}
