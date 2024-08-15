package app

import (
	"{appname}/internal/repository"
)

type demoApp struct {
	*AppBasic
	demoRepository repository.DemoRepository
}

func NewDemoApp(appimp *AppBasic, demoRepository repository.DemoRepository) App {
	return &demoApp{
		AppBasic:       appimp,
		demoRepository: demoRepository,
	}
}

func (ua *demoApp) Run() {
	ua.logger.Info("app started")
}
