package database

import "gorm.io/gorm"

const (
	SQL              = "sql"
	NoSQL            = "nosql"
	DriverPostgresql = "postgresql"
	DriverSqlite     = "sqlite"
)

type Config struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Options  string
}

type Database struct {
	SQL map[string]*gorm.DB
}

func NewDatabase() *Database {
	return &Database{
		SQL: make(map[string]*gorm.DB),
	}
}

func (d Database) GetSQL(name string) *gorm.DB {
	db, ok := d.SQL[name]
	if !ok {
		return nil
	}

	return db
}

func (d *Database) AddSQL(name string, config Config) error {
	var db *gorm.DB
	var err error
	if config.Driver == DriverPostgresql {
		db, err = ConfigurePostgresql(config)
	} else if config.Driver == DriverSqlite {
		db, err = ConfigureSqlite()
	}

	if err != nil {
		return err
	}

	d.SQL[name] = db
	return nil
}
