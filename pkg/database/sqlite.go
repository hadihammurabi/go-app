package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConfigureSqlite func
func ConfigureSqlite() (*gorm.DB, error) {
	location := os.Getenv("DB_LOCATION")
	if location == "" {
		location = "db.sqlite3"
	}

	db, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	return db, err
}
