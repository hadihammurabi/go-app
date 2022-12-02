package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Options  string
}

// ConfigureMysql func
func ConfigureMysql(config Config) (*gorm.DB, error) {
	host := config.Host
	if host == "" {
		host = "localhost"
	}

	port := fmt.Sprintf("%d", config.Port)
	if port == "" {
		port = "3306"
	}

	user := config.Username
	if user == "" {
		user = "root"
	}

	pass := config.Password
	name := config.Name

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		user,
		pass,
		host,
		port,
		name,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}
