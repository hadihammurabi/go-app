package database

import (
	"errors"

	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func configureSql(conf config.Database) error {
	var err error
	if conf.Driver == "postgresql" {
		err = configurePostgresql(conf)
	} else if conf.Driver == "mysql" {
		err = configureMysql(conf)
	} else if conf.Driver == "sqlite" {
		err = configureSqlite(conf)
	} else {
		err = errors.New("unknown database driver")
	}

	return err
}

func configureMysql(conf config.Database) error {
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	ioc.Set(func() database.MySQL { return database.MySQL{DB: db} })
	return nil
}

func configurePostgresql(conf config.Database) error {
	db, err := gorm.Open(postgres.Open(conf.DSN), &gorm.Config{})
	if err != nil {
		return err
	}

	ioc.Set(func() database.PostgreSQL { return database.PostgreSQL{DB: db} })
	return nil
}

func configureSqlite(conf config.Database) error {
	location := conf.DSN
	if location == "" {
		location = "db.sqlite3"
	}

	db, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	if err != nil {
		return err
	}

	ioc.Set(func() database.SQLite { return database.SQLite{DB: db} })
	return nil
}
