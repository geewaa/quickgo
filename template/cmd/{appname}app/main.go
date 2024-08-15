package main

import (
	"{appname}/cmd/{appname}app/wire"
	"{appname}/pkg/config"
	"{appname}/pkg/log"
)

//go:generate wire ./wire
func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	app, cleanup, err := wire.NewWire(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Run()
}
