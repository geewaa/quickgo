package app

import "{appname}/pkg/log"

type App interface {
	Run()
}

type AppBasic struct {
	logger *log.Logger
}

func NewAppBasic(logger *log.Logger) *AppBasic {
	return &AppBasic{
		logger: logger,
	}
}
