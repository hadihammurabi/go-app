package database

import (
	"github.com/hadihammurabi/belajar-go-rest-api/driver/database/sql"
	"gorm.io/gorm"
)

type Config struct {
	Driver   DatabaseDriver
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Options  string
}

type Database struct {
	*gorm.DB
	Connections map[string]*gorm.DB
}

func NewDatabase() *Database {
	return &Database{
		Connections: make(map[string]*gorm.DB),
	}
}

func (d Database) GetConnection(names ...string) *gorm.DB {
	name := ""
	if len(names) > 0 {
		name = names[0]
	}

	if name == "" {
		return d.DB
	}

	db, ok := d.Connections[name]
	if !ok {
		return nil
	}

	return db
}

func (d *Database) AddConnection(name string, config Config) error {
	var db *gorm.DB
	var err error
	if config.Driver.Driver == sql.DriverPostgresql {
		db, err = sql.ConfigurePostgresql(sql.Config{
			Host:     config.Host,
			Port:     config.Port,
			Username: config.Username,
			Password: config.Password,
			Name:     config.Name,
			Options:  config.Options,
		})
	} else if config.Driver.Driver == sql.DriverMysql {
		db, err = sql.ConfigureMysql(sql.Config{
			Host:     config.Host,
			Port:     config.Port,
			Username: config.Username,
			Password: config.Password,
			Name:     config.Name,
			Options:  config.Options,
		})
	} else if config.Driver.Driver == sql.DriverSqlite {
		db, err = sql.ConfigureSqlite()
	}

	if err != nil {
		return err
	}

	d.Connections[name] = db
	return nil
}
