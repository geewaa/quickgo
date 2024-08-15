//go:build wireinject
// +build wireinject

package wire

import (
	"{appname}/internal/app"
	"{appname}/internal/repository"
	"{appname}/pkg/log"

	"github.com/google/wire"
	"github.com/spf13/viper"
)

var RepositorySet = wire.NewSet(
	repository.NewDb,
	repository.NewRepository,
	repository.NewDemoRepository,
)

var AppSet = wire.NewSet(
	app.NewAppBasic,
	app.NewDemoApp,
)

func NewWire(*viper.Viper, *log.Logger) (app.App, func(), error) {
	panic(wire.Build(
		RepositorySet,
		AppSet,
	))
}
