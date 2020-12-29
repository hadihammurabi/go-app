package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConfigureSqlite func
func ConfigureSqlite(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Location), &gorm.Config{})
	return db, err
}
