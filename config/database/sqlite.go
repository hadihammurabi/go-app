package database

import (
	"belajar-go-rest-api/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConfigureSqlite func
func ConfigureSqlite(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Location), &gorm.Config{})
	return db, err
}

func migrateSqlite(db *gorm.DB) {
	db.AutoMigrate(entity.User{})
}
