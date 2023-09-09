package driver

import (
	"context"
	"os"
	"testing"

	"github.com/golang-must/must"
	"github.com/gowok/gowok/exception"
	"github.com/hadihammurabi/belajar-go-rest-api/pkg"
)

func TestDriverGet(t *testing.T) {
	defer func() {
		err := recover()
		must.Nil(t, err)
	}()
	confPath, err := pkg.GetConfigPath("..")
	must.Nil(t, err)

	os.Args = []string{"", "-config", confPath}
	d := Get()

	sqlO := 0
	err = d.SQL.Get().OrPanic(exception.ErrNoDatabaseFound).Raw("SELECT 1").Scan(&sqlO).Error
	must.Nil(t, err)
	must.Equal(t, sqlO, 1)

	err = d.MongoDB.Get("secondary").OrPanic(exception.ErrNoDatabaseFound).Ping(context.Background(), nil)
	must.Nil(t, err)

	err = d.Redis.Get("cache").OrPanic(exception.ErrNoDatabaseFound).Ping(context.Background()).Err()
	must.Nil(t, err)
}
