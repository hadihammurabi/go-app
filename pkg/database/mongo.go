package database

import (
	"github.com/gowok/gowok/config"
	"github.com/gowok/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func configureMongo(conf config.Database) error {
	client, err := mongo.NewClient(options.Client().ApplyURI(conf.DSN))
	if err != nil {
		return err
	}

	ioc.Set(func() mongo.Client { return *client })

	return err
}
