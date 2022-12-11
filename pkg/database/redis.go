package database

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gowok/gowok/config"
	"github.com/gowok/gowok/driver/database"
	"github.com/gowok/ioc"
)

func configureRedis(conf config.Database) error {
	client := redis.NewClient(&redis.Options{
		Addr: conf.DSN,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	ioc.Set(func() database.Redis { return database.Redis{Client: client} })

	return err
}
