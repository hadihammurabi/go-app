package config

import (
	"belajar-go-rest-api/model"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConfigureDatabase func
func ConfigureDatabase() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			port,
		),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	return db, err
}

// MigrateDatabase func
func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
	)
}
