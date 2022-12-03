package sql

type Config struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Options  string
}

const (
	DriverPostgresql = "postgresql"
	DriverSqlite     = "sqlite"
	DriverMysql      = "mysql"
)
