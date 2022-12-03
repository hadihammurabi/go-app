package sql

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConfigureSqlite func
func ConfigureSqlite(config Config) (*gorm.DB, error) {
	location := config.DSN
	if location == "" {
		location = "db.sqlite3"
	}

	db, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	return db, err
}
