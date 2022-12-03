package config

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/hadihammurabi/belajar-go-rest-api/driver/database"
	"github.com/spf13/viper"
)

type dbconfig struct {
	ID       string
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Options  string
}

func (d dbconfig) ToDBConfig() (database.Config, bool) {
	driver, ok := database.MapDriver(d.Driver)
	if !ok {
		return database.Config{}, false
	}

	return database.Config{
		Driver:   driver,
		Host:     d.Host,
		Port:     d.Port,
		Username: d.Username,
		Password: d.Password,
		Name:     d.Name,
		Options:  d.Options,
	}, true
}

func dbconfigFromMap(id string, in map[string]any) dbconfig {
	in["id"] = id
	inJSON, _ := json.Marshal(in)

	config := &dbconfig{}
	json.Unmarshal(inJSON, config)

	return *config
}

// ConfigureDatabase func
func ConfigureDatabase() *database.Database {
	dbconfigsFromConfig := viper.Get("databases").(map[string]any)

	if dbconfigsFromConfig[dbconfigsFromConfig["active"].(string)] == nil {
		panic(fmt.Sprintf("database %s not configured", dbconfigsFromConfig["active"]))
	}

	dbconfigs := []dbconfig{}
	for id, configs := range dbconfigsFromConfig {
		if id != "active" {
			dbconfigs = append(dbconfigs, dbconfigFromMap(id, configs.(map[string]any)))
		}
	}

	db := database.NewDatabase()
	for _, config := range dbconfigs {
		dbconfig, ok := config.ToDBConfig()
		if !ok {
			log.Println("database configuration failed")
		}

		err := db.AddConnection(config.ID, dbconfig)
		if err != nil {
			log.Println(err)
		}
	}
	db.DB = db.GetConnection(dbconfigsFromConfig["active"].(string))

	return db
}
