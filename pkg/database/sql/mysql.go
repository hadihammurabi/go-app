package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConfigureMysql func
func ConfigureMysql(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})

	return db, err
}
