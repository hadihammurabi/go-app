package sql

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConfigurePostgresql func
func ConfigurePostgresql(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})

	return db, err
}
