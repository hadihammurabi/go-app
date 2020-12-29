package database

import (
	"belajar-go-rest-api/entity"
	"errors"
	"os"

	"gorm.io/gorm"
)

const (
	driverPostgresql = "postgresql"
	driverSqlite     = "sqlite"
)

// DBConfig struct
type DBConfig struct {
	Driver string
	Host   string
	Port   string

	Location string

	User string
	Pass string
	Name string
}

// ConfigureDatabase func
func ConfigureDatabase() (*gorm.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	if driver == "" {
		driver = driverPostgresql
	}

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	location := os.Getenv("DB_LOCATION")
	if location == "" {
		location = "db.sqlite3"
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	config := &DBConfig{
		Driver:   driver,
		Host:     host,
		Port:     port,
		Location: location,
		User:     user,
		Pass:     pass,
		Name:     name,
	}

	if driver == driverPostgresql {
		db, err := ConfigurePostgresql(config)
		return db, err
	} else if driver == driverSqlite {
		db, err := ConfigureSqlite(config)
		migrateSqlite(db)
		return db, err
	}

	return nil, errors.New("Unknown database driver")
}

// MigrateDatabase func
func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&entity.User{},
	)
}
