package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConfigurePostgresql func
func ConfigurePostgresql(config *DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Host,
			config.User,
			config.Pass,
			config.Name,
			config.Port,
		),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	return db, err
}
