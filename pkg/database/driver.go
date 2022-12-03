package database

import (
	"golang.org/x/exp/slices"
)

type DBType string

const (
	SQL   DBType = "sql"
	NoSQL DBType = "nosql"
)

type DatabaseDriver struct {
	Type   DBType
	Driver string
}

var sqlDrivers = []string{"mysql", "mariadb", "postgresql", "sqlite"}
var mongoDrivers = []string{"mongodb"}

func MapDriver(name string) (DatabaseDriver, bool) {
	driver := DatabaseDriver{Driver: name}

	isSQL := slices.Contains(sqlDrivers, driver.Driver)
	if isSQL {
		driver.Type = SQL
		return driver, true
	}

	isMongo := slices.Contains(mongoDrivers, driver.Driver)
	if isMongo {
		driver.Type = SQL
		return driver, true
	}

	return DatabaseDriver{}, false
}
