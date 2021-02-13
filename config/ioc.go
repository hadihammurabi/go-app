package config

import (
	"github.com/hadihammurabi/belajar-go-rest-api/repository"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

// NewIOC func
func NewIOC(db *gorm.DB) di.Container {
	builder, _ := di.NewBuilder()
	builder.Add(di.Def{
		Name: "database",
		Build: func(ctn di.Container) (interface{}, error) {
			return db, nil
		},
	})

	builder.Add(di.Def{
		Name: "repository",
		Build: func(ctn di.Container) (interface{}, error) {
			return repository.NewRepository(builder.Build()), nil
		},
	})

	return builder.Build()
}
