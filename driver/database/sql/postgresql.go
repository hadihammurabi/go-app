package sql

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConfigurePostgresql func
func ConfigurePostgresql(config Config) (*gorm.DB, error) {
	host := config.Host
	if host == "" {
		host = "localhost"
	}

	port := fmt.Sprintf("%d", config.Port)
	if port == "" {
		port = "5432"
	}

	user := config.Username
	if user == "" {
		user = "postgres"
	}

	pass := config.Password
	name := config.Name

	dsn := fmt.Sprintf(
		"host=%s user=%s database=%s port=%s sslmode=disable",
		host,
		user,
		name,
		port,
	)
	if pass != "" {
		dsn += " password=" + pass
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
