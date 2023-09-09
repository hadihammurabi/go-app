package service

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/golang-must/must"
	"github.com/google/uuid"
	"github.com/gowok/gowok/exception"
	"github.com/hadihammurabi/belajar-go-rest-api/driver"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestServiceUser(t *testing.T) {
	confPath, err := pkg.GetConfigPath("..")
	must.Nil(t, err)

	os.Args = []string{"", "-config", confPath}
	dr := driver.Get()
	repo := repository.Get()
	sqldb := dr.SQL.Get().OrPanic(exception.ErrNoDatabaseFound)
	sqldb.Logger = logger.Default.LogMode(logger.Silent)
	sv := NewUserService(dr.Config, sqldb, repo)

	t.Run("Get All", func(t *testing.T) {
		_, err := sv.All(context.Background())
		must.Nil(t, err)
	})

	t.Run("Get By ID", func(t *testing.T) {
		_, err := sv.FindByID(context.Background(), uuid.New())
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			must.Nil(t, err)
		}
	})
}
