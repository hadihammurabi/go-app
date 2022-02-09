package config

import (
	"log"

	"github.com/hadihammurabi/belajar-go-rest-api/pkg/database"
	"github.com/spf13/viper"
)

type dbconfig struct {
	ID       string
	Type     string
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Panic    bool
	Options  string
}

func (d dbconfig) ToDBConfig() database.Config {
	return database.Config{
		Driver:   d.Driver,
		Host:     d.Host,
		Port:     d.Port,
		Username: d.Username,
		Password: d.Password,
		Name:     d.Name,
		Options:  d.Options,
	}
}

// ConfigureDatabase func
func ConfigureDatabase() *database.Database {
	dbconfigsFromConfig := viper.Get("databases").([]interface{})
	dbconfigs := []dbconfig{}
	for _, db := range dbconfigsFromConfig {
		dbmap := db.(map[interface{}]interface{})
		if dbmap["id"] == nil {
			dbmap["id"] = ""
		}
		if dbmap["type"] == nil {
			dbmap["type"] = ""
		}
		if dbmap["driver"] == nil {
			dbmap["driver"] = ""
		}
		if dbmap["host"] == nil {
			dbmap["host"] = ""
		}
		if dbmap["port"] == nil {
			dbmap["port"] = 0
		}
		if dbmap["username"] == nil {
			dbmap["username"] = ""
		}
		if dbmap["password"] == nil {
			dbmap["password"] = ""
		}
		if dbmap["name"] == nil {
			dbmap["name"] = ""
		}
		if dbmap["panic"] == nil {
			dbmap["panic"] = false
		}
		if dbmap["options"] == nil {
			dbmap["options"] = ""
		}

		dbconfigs = append(dbconfigs, dbconfig{
			ID:       dbmap["id"].(string),
			Name:     dbmap["name"].(string),
			Type:     dbmap["type"].(string),
			Driver:   dbmap["driver"].(string),
			Host:     dbmap["host"].(string),
			Port:     dbmap["port"].(int),
			Username: dbmap["username"].(string),
			Password: dbmap["password"].(string),
			Options:  dbmap["options"].(string),
			Panic:    dbmap["panic"].(bool),
		})
	}

	db := database.NewDatabase()
	for _, config := range dbconfigs {
		if config.Type == database.SQL {
			err := db.AddSQL(config.ID, config.ToDBConfig())
			if err != nil {
				switch config.Panic {
				case true:
					panic(err)
				case false:
					log.Println(err)
				}
			}
		}
	}

	return db
}
